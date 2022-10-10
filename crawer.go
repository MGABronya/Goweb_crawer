// @Title  crawer
// @Description  设计一个简单的爬虫
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-10-10 00:39
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// @title    Get
// @description   获取一个url页面的内容
// @auth      MGAronya（张健）             2022-10-10 00:39
// @param     url string				要被提取的url页面
// @return    result string, err error			result表示提取出来的信息，err表示潜在的报错
func Get(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4*1024)
	// TODO 读取网站的Body内容
	for {
		n, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕")
				break
			} else {
				fmt.Println("resp.Body.Read err =", err)
				break
			}
		}
		result += string(buf[:n])
	}
	return
}

// @title    SpiderPage
// @description   将页面内容写入对应的文件里
// @auth      MGAronya（张健）             2022-10-10 00:39
// @param     i int, page chan<- int	  i表示要爬取是第几页，page用来传递爬取消息
// @return    void
func SpiderPage(i int, page chan<- int) {
	url := "https://github.com/search?q=go&type=Repositories&p=1" + strconv.Itoa((i-1)*50)
	fmt.Printf("正在爬取第%d个网页\n", i)
	// TODO 爬取网页的信息
	result, err := Get(url)
	if err != nil {
		fmt.Println("http.Get err = ", err)
		return
	}
	// TODO 将内容写入文件
	filename := "page" + strconv.Itoa(i) + ".html"
	f, err1 := os.Create(filename)
	if err1 != nil {
		fmt.Println("os.Create err =", err1)
		return
	}
	// TODO 写入内容
	f.WriteString(result)
	// TODO 关闭文件
	f.Close()
	page <- i
}

// @title    Run
// @description   让每个页面都单独运行一个goroutine
// @auth      MGAronya（张健）             2022-10-10 00:39
// @param     start, end int		表示起始页和结束页
// @return    void
func Run(start, end int) {
	fmt.Printf("正在爬取第%d页到%d页\n", start, end)

	// TODO 需要将数据传入通道
	page := make(chan int)
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}

	// TODO 等待所有爬取完成
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}
}

// @title    main
// @description   运行整个项目
// @auth      MGAronya（张健）             2022-10-10 00:39
// @param     void
// @return    void
func main() {
	var start, end int
	fmt.Printf("请输入起始页>=1: > ")
	fmt.Scan(&start)
	fmt.Printf("请输入结束页数字: ? ")
	fmt.Scan(&end)
	Run(start, end)
}
