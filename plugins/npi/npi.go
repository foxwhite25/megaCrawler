package npi

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("npi", "世界和平研究所", "https://npi.or.jp/")
	w.SetStartingUrls([]string{"https://npi.or.jp/research/diplomacy/index.html",
		"https://npi.or.jp/research/industry/index.html",
		"https://npi.or.jp/research/economy/index.html",
		"https://npi.or.jp/research/technology/index.html",
		"https://npi.or.jp/research/government/index.html",
		"https://npi.or.jp/research/policy/index.html",
		"https://npi.or.jp/experts/index.html"})

	//index
	w.OnHTML("a.link_page", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//report
	w.OnHTML("#main > div.entry_list_area > div > div> p > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	//标题
	w.OnHTML("h3.ttl_second.b30>div", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//时间
	w.OnHTML(" h3 > div.date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//正文
	w.OnHTML("#main > div.detail_area > div.b0 > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	//pdf
	w.OnHTML("a.btn_pdf", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})

	//访问专家
	w.OnHTML("a.btn_bluecorner", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//姓名
	w.OnHTML("div.name.b05", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	//头衔
	w.OnHTML("#main > div:nth-child(1) > div > h3 > div.date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//领域
	w.OnHTML("#main > p:nth-child(5)", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area = element.Text
	})
}
