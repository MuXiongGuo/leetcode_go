package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 网页规律
// 1. page的规律 页面变化的规律: https://duanzixing.com/page/3/
// 2. 当前页面的规律(配合源码 + ctrl f 来寻找规律): <h2><a href="url"
// 3. 找到标题的规律 <title>xxx
// 4. 提取到标题并进入页面中
// 5. 找到标题是哪个  内容是哪个, 可能要考虑是第几个的问题
// 6. content <article class="article-content">(?s:(.*?))</p>                </article>

func ReMatch(formula, result string) (data []string) {
	// 匹配
	re := regexp.MustCompile(formula)
	if re == nil {
		fmt.Println("re error!")
	}
	objUrls := re.FindAllStringSubmatch(result, -1)
	// objUrls 数组的[1]就是自动去掉头尾的, [0]是带头尾的, 用起来非常方便
	for _, val := range objUrls {
		data = append(data, val[1])
	}
	return
}

func SpiderOneJoy(url string) (title, content string, err error) {
	result, err := HttpGet(url)
	if err != nil {
		return
	}
	// 取标题
	if len(ReMatch(`<title>(?s:(.*?)) `, result)) == 0 || len(ReMatch(`<p>(?s:(.*?))</p>                </article>`, result)) == 0 {
		fmt.Println(url, "Bad 502")
		content = "502 error"
		title = "502 error"
		return
	}
	title = ReMatch(`<title>(?s:(.*?)) `, result)[0]
	// 取内容
	content = ReMatch(`<p>(?s:(.*?))</p>                </article>`, result)[0]
	content = strings.Replace(content, "<br>", "\n", -1)
	content = strings.Replace(content, "</p><p>", "", -1)
	return
}

func HttpGet(url string) (result string, err error) { // 读url里面的body内容
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 读取网页内容
	buf := make([]byte, 4*1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf[:n])
	}
	return
}

func SpiderPage(i int, page chan int) {
	url := "https://duanzixing.com/page/" + strconv.Itoa(i) + "/"
	fmt.Println("Crawling page " + strconv.Itoa(i) + ":" + url)

	// 主页面爬取 取链接
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	// ?s: 单行模式, .*?懒惰模式匹配所有可能
	joyUrls := ReMatch(`<h2><a href="(?s:(.*?))"`, result)
	// 进入网址里面继续爬
	titlePage, contentPage := make([]string, 0), make([]string, 0)
	for _, data := range joyUrls {
		title, content, err := SpiderOneJoy(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		titlePage = append(titlePage, title)
		contentPage = append(contentPage, content)
	}
	f, err := os.Create(strconv.Itoa(i) + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(titlePage); i++ {
		f.WriteString("title" + strconv.Itoa(i) + ": " + titlePage[i] + "\n")
		f.WriteString("content" + strconv.Itoa(i) + ": " + contentPage[i] + "\n\n")
	}
	page <- i
}

func DoWork(start, end int) {
	fmt.Printf("starting get page %d to page %d\n", start, end)
	page := make(chan int, 1)
	for i := start; i <= end; i++ {
		// 爬主页面
		go SpiderPage(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Println("page " + strconv.Itoa(<-page) + "complete")
	}
}

func main() {
	var start, end int
	fmt.Printf("input start page(>=1): ")
	fmt.Scan(&start)
	fmt.Printf("input end page(>=start page): ")
	fmt.Scan(&end)
	startTime := time.Now()
	DoWork(start, end)
	costTime := time.Since(startTime)
	fmt.Println("cost time:", costTime)
}
