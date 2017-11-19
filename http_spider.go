package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//先爬取每个菜单列表主页， 然后再根据菜单列表 进入每个段子的主页，将段子的内容和标题拿到
//菜单页的规则
//第1页 http://www.neihan8.com/article/index.html
//第2页 http://www.neihan8.com/article/index_2.html
//第3页 http://www.neihan8.com/article/index_3.html
//第n页 http://www.neihan8.com/article/index_n.html

//从菜单页得到每个段子的跳转路径的正则规则
//`<h3><a href="(?s:(.*?))"`  ---> /article/209244.html ---> http://www.neihan8.com/article/209244.html

//再去进入爬取段子信息
//标题的正则规则
//`<h1 class="title">(?s:(.*?))</h1>`
//段子内容的正则规则
//`</a></p>(?s:(.*?))<div class="ad610">`

type Spider struct {
	page int
}

func (this *Spider) HttpGet(url string) (content string, statusCode int) {
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

//把段子存储到文件中
func (this *Spider) Store_dz_to_file(title []string, contents []string) {
	filename := "MyduanZi.txt"

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file error")
		return
	}
	defer f.Close()

	for i := 0; i < len(title); i++ {
		f.WriteString("\n-------------------\n")
		f.WriteString(title[i])
		f.WriteString("\n-------------------\n")
		f.WriteString(contents[i])
	}
}

func (this *Spider) SpiderOneDuanz(url string) (dz_title string, dz_content string) {
	fmt.Println("正在爬取 ", url)
	content, rcode := this.HttpGet(url)
	if rcode < 0 {
		fmt.Println("http get error rcode = ", rcode)
		return "", ""
	}

	//得到标题
	title := regexp.MustCompile(`<h1 class="title">(?s:(.*?))</h1>`)
	titles := title.FindAllStringSubmatch(content, -1)
	for _, Title := range titles {
		dz_title = Title[1]
		break
	}
	//得到内容
	content_exp := regexp.MustCompile(`</a></p>(?s:(.*?))<div class="ad610">`)
	contents := content_exp.FindAllStringSubmatch(content, -1)
	for _, content_dz := range contents {
		dz_content = content_dz[1]
		dz_content = strings.Replace(dz_content, "\r\n", "\n", -1)
		dz_content = strings.Replace(dz_content, "<p>", "", -1)
		dz_content = strings.Replace(dz_content, "</p>", "", -1)
		dz_content = strings.Replace(dz_content, "&ldquo", "", -1)
		dz_content = strings.Replace(dz_content, "&hellip", "", -1)
		dz_content = strings.Replace(dz_content, "&rdquo", "", -1)
		dz_content = strings.Replace(dz_content, "&nbsp", "", -1)
		dz_content = strings.Replace(dz_content, ";", "", -1)
		break
	}
	return
}

func (this *Spider) SpiderOnePage() {
	fmt.Println("正在爬取 第", this.page, "页")
	url := ""
	if this.page == 1 {
		url = "http://www.neihan8.com/article/index.html"
	} else {
		url = "http://www.neihan8.com/article/index_" + strconv.Itoa(this.page) + ".html"
	}
	fmt.Println("url = ", url)

	content, rcode := this.HttpGet(url)
	if rcode < 0 {
		fmt.Println("http get error rcode = ", rcode)
		return
	}

	//当前页面的段子标题和内容

	title_slice := make([]string, 0)
	content_slice := make([]string, 0)

	dz_url := regexp.MustCompile(`<h3><a href="(?s:(.*?))"`)
	urls := dz_url.FindAllStringSubmatch(content, -1)
	for _, dz_url := range urls {
		full_url := "http://www.neihan8.com" + dz_url[1]
		//爬取一个段子
		title, content := this.SpiderOneDuanz(full_url)
		//fmt.Println(dz_url[1])
		title_slice = append(title_slice, title)
		content_slice = append(content_slice, content)
	}
	//把当前页面的爬取到的全部段子存入到文件中
	this.Store_dz_to_file(title_slice, content_slice)

}

func (this *Spider) Dowork() {
	fmt.Println("Spider begin to work ")
	this.page = 1

	var cmd string

	for {
		fmt.Println("请输入任意键 爬取下一页，输入exit 退出")
		fmt.Scanf("%s", &cmd)

		if cmd == "exit " {
			fmt.Println("exit")
			break
		}
		this.SpiderOnePage()

		this.page++
	}
}
func main() {

	sp := new(Spider)
	sp.Dowork()

}
