package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

// 网页规律
// 1. page的规律 页面变化的规律: https://duanzixing.com/page/3/
// 2. 当前页面的规律(配合源码 + ctrl f 来寻找规律): <h2><a href="url"
// 3. 找到标题的规律 <title>xxx
// 4. 提取到标题并进入页面中
// 5. 找到标题是哪个  内容是哪个, 可能要考虑是第几个的问题

func HttpGet(url string) (result string, err error) {
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

func SpiderPage(i int) {
	url := "https://duanzixing.com/page/" + strconv.Itoa(i) + "/"
	fmt.Println("Crawling page " + strconv.Itoa(i) + ":" + url)

	// 主页面爬取 取链接
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	// ?s: 单行模式, .*?懒惰模式匹配所有可能
	re := regexp.MustCompile(`<h2><a href="(?s:(.*?))"`)
	if re == nil {
		fmt.Println("re error!")
	}
	joyUrls := re.FindAllStringSubmatch(result, -1)
	// joyUrls 数组的[1]就是自动去掉头尾的, [0]是带头尾的, 用起来非常方便
	// 进入网址里面继续爬
	for _, data := range joyUrls {
		title, content, err := SpiderOneJoy(data[1])
	}
}

func SpiderOneJoy(url string) (title, content string, err error) {
	result, err := HttpGet(url)
	if err != nil {
		return
	}
	// 取标题<h1 >
}

func DoWork(start, end int) {
	fmt.Printf("starting get page %d to page %d\n", start, end)

	for i := start; i <= end; i++ {
		// 爬主页面
		SpiderPage(i)
	}
}

func main() {
	var start, end int
	fmt.Printf("input start page(>=1): ")
	fmt.Scan(&start)
	fmt.Printf("input end page(>=start page): ")
	fmt.Scan(&end)
	DoWork(start, end)
}
