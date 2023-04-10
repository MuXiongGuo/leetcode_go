package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// 爬取百度贴吧
// https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=0
// pn=0 第一页 50 第二页 以此类推

func HttpGet(url string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 读取网页内容(body内容)
	buf := make([]byte, 1024*4)
	for {
		n, err1 := resp.Body.Read(buf)
		if n == 0 {
			//fmt.Println("Body: ", err1)
			err = err1
			break
		}
		result += string(buf[:n]) // 推荐用stringBuilder而不是+=
	}
	return
}

func Dowork(start, end int) {
	fmt.Printf("getting the content of %d to %d\n", start, end)
	page := make(chan int)
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}
	// 注意不能让主程序提前结束
	for i := start; i <= end; i++ {
		fmt.Printf("page %d complete \n", <-page)
	}
}

func SpiderPage(i int, page chan int) {
	// 1.爬取目标
	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" +
		strconv.Itoa((i-1)*50)
	fmt.Println("crawling page " + strconv.Itoa(i))
	// 2.爬取
	result, err := HttpGet(url)
	if err != nil && err.Error() != "EOF" {
		fmt.Println(err)
		return
	}
	// 3.把内容写入到文件
	fileName := strconv.Itoa(i) + ".html"
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println("os.Create err: ", err)
		return
	}
	f.WriteString(result) // 写入内容
	page <- i
}
func main() {
	//runtime.GOMAXPROCS(1)

	var start, end int
	fmt.Printf("input start page(>=1): ")
	fmt.Scan(&start)
	fmt.Printf("input end page(>=start page): ")
	fmt.Scan(&end)
	Dowork(start, end)
}
