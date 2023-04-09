package main

import (
	"fmt"
	"strconv"
)

// 爬取百度贴吧
// https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=0
// pn=0 第一页 50 第二页 以此类推

func HttpGet(url string) {

}

func Dowork(start, end int) {
	fmt.Printf("getting the content of %d to %d", start, end)

	// 1.爬取目标
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" +
			strconv.Itoa((i-1)*50)
		// 2.爬取
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		SpiderOnePage()
	}
}
func main() {
	var start, end int
	fmt.Printf("input start page(>=1): ")
	fmt.Scan(&start)
	fmt.Printf("input end page(>=start page): ")
	fmt.Scan(&end)
	Dowork(start, end)
}
