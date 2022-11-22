package zri

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("zri", "Zavecz Research Institute", "http://www.zaveczresearch.hu/")
	w.SetStartingUrls([]string{"https://www.zaveczresearch.hu/category/elemzesek/",
		"http://www.zaveczresearch.hu/kutatasi-eredmenyeink/"})

	//index
	w.OnHTML("div.nav-previous > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问新闻
	w.OnHTML("h3.entry-title > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//标题
	w.OnHTML("h1.entry-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//正文
	w.OnHTML("div.entry-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
}
