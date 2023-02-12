package aseantribune

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("aseantribune", "东盟论坛报", "https://www.aseantribune.com/")

	w.SetStartingUrls([]string{"https://www.aseantribune.com/category/general/",
		"https://www.aseantribune.com/category/trading/",
		"https://www.aseantribune.com/category/science/",
		"https://www.aseantribune.com/category/self-care/",
		"https://www.aseantribune.com/category/key-issues/",
		"https://www.aseantribune.com/category/legal-matters/",
		"https://www.aseantribune.com/category/press-releases/"})

	// 从翻页器获取链接并访问
	w.OnHTML("div.nav-previous>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从index访问新闻
	w.OnHTML("h2.entry-title>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	// report.title
	w.OnHTML("h1.entry-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// report .content
	w.OnHTML("div.entry-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	//report.publish time
	w.OnHTML("time.entry-date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.author
	w.OnHTML("div.entry-meta>span.byline>span.author>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

}
