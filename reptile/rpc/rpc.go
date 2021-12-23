package main

import (
	"flag"
	"fmt"
	"novel/projectUtil"

	"github.com/gocolly/colly"
)

// var configFile = flag.String("f", "etc/rpc.yaml", "the config file")

var log = projectUtil.GetFileLog()

func main() {
	flag.Parse()

	// 加载配置文件
	// var c config.Config
	// conf.MustLoad(*configFile, &c)
	// ctx := svc.NewServiceContext(c)

	// logx.Info("hello")

	// log.Info(GetMenu())
	GetMenu()
}

// GetMenu 获取到对应的小说名和链接
func GetMenu() (data string) {
	c := colly.NewCollector()

	c.OnHTML("div[id='main'] ul li a", func(e1 *colly.HTMLElement) {
		fmt.Println(e1.Text, e1.Attr("href"))
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("visiting", request.URL)
	})

	c.Visit("https://www.xbiquge.la/xiaoshuodaquan/")

	return ""
}

// GetDetail 获取到对应的详情和章节列表
func GetDetail(url string) {
	//c := colly.NewCollector()
	//
	//// 获取小说名字，小说类型，小说作者，最后更新时间
	//c.OnHTML("div[id='bdshare']", func(e1 *colly.HTMLElement) {
	//	e1.ForEach("div > div[id=""]")
	//})
	//
	//// 获取章节列表
	//c.OnHTML("div[id='main'] ul li a", func(e1 *colly.HTMLElement) {
	//	fmt.Println(e1.Text, e1.Attr("href"))
	//})
	//
	//c.OnRequest(func(request *colly.Request) {
	//	fmt.Println("visiting detail:", request.URL)
	//})
	//
	//c.Visit(url)

	return
}

// GetContent 获取到对应的章节内容
func GetContent(url string) {

}
