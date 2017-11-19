package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

//第1页：https://tieba.baidu.com/f?kw=lol&ie=utf-8&pn=0
//第2页：https://tieba.baidu.com/f?kw=lol&ie=utf-8&pn=50
//第3页：https://tieba.baidu.com/f?kw=lol&ie=utf-8&pn=100
//第4页：https://tieba.baidu.com/f?kw=lol&ie=utf-8&pn=150

func httpGet(url string) (content string, statusCode int) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		statusCode = -100
		return
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		statusCode = resp.StatusCode
		return
	}

	content = string(data)
	statusCode = resp.StatusCode
	return

}

func spider_tieba(begin int, end int) {

	var pn int
	fmt.Println("准备爬取 从 ", begin, " 到 ", end, "页")

	for page := begin; page < end+1; page++ {
		fmt.Println("正在爬取 第", page, "页")

		pn = (page - 1) * 50
		url := "https://tieba.baidu.com/f?kw=lol&ie=utf-8&pn= " + strconv.Itoa(pn)
		fmt.Println("url = ", url)

		//开始爬取
		content, rcode := httpGet(url)

		if rcode < 0 {
			fmt.Println("httpGet error,rcode = ", rcode, "page = ", page)
			continue
		}

		//处理数据
		filename := strconv.Itoa(page) + ".html"
		if f, err := os.Create(filename); err == nil {
			f.WriteString(content)
			f.Close()
		}
	}

}

func main() {
	fmt.Println("vim-go")
	fmt.Println("====================")

	var begin string
	var end string

	fmt.Println("请输入要爬取的起始页码")
	fmt.Scanf("%s\n", &begin)
	fmt.Println("请输入要爬取的终止页码")
	fmt.Scanf("%s\n", &end)

	b, _ := strconv.Atoi(begin)
	e, _ := strconv.Atoi(end)

	spider_tieba(b, e)

}
