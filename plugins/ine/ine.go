package ine

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func cutToList(inputStr string) []string {
	nameList := strings.Split(inputStr, ",")
	for index, value := range nameList {
		nameList[index] = strings.TrimSpace(value)
	}

	return nameList
}

func init() {
	w := Crawler.Register("ine", "西班牙国家统计局", "https://www.ine.es/")
	w.SetStartingUrls([]string{"https://www.ine.es/ss/Satellite?L=en_GB&c=Page&cid=1254735839320&p=1254735839320&pagename=MetodologiaYEstandares%2FINELayout"})

	//访问报告
	w.OnHTML("table > tbody > tr> td> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	//报告标题
	w.OnHTML("h1.revista_tit", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//pdf
	w.OnHTML("div.cuerpo_principal > div> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})

	//作者
	w.OnHTML(" div.cuerpo_principal > div> em", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = cutToList(element.Text)
	})

	//正文
	w.OnHTML("div.cuerpo_principal", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
}
