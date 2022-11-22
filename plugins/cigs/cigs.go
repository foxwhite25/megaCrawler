package cigs

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("cigs", "佳能国际问题研究中心", "https://cigs.canon/en/")
	w.SetStartingUrls([]string{"https://cigs.canon/en/fellows/"})

	//访问专家
	w.OnHTML("#nav-ja-order > div > div> div.card-fellows > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//获取姓名
	w.OnHTML("p.name_text", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	//获取头衔
	w.OnHTML("p.position_text", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//获取领域
	w.OnHTML("#profile1 > div > div.col-12.col-md-12.pt-5.pb-5 > section > ul > li", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area = element.Text
	})

	//访问报告
	w.OnHTML("div > div > div > section > div > div > div > div > h3 > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	//pdf
	w.OnHTML(" div > div> section> p > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})
	w.OnHTML("p.pdf-link > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})
	w.OnHTML("a.newwin.iconPDF", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})

	//标题
	w.OnHTML("#contents > div> div > div > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//正文
	w.OnHTML(" section > div > div > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML(" section > div > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content += element.Text
	})
}
