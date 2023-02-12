package gna

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("gn", "印度政策研究中心", "https://gna.org.gh/")

	w.SetStartingUrls([]string{"https://gna.org.gh/headline/",
		"https://gna.org.gh/business/",
		"https://gna.org.gh/ghana-economic-data/",
		"https://gna.org.gh/politics/",
		"https://gna.org.gh/education/",
		"https://gna.org.gh/health/",
		"https://gna.org.gh/crime/",
		"https://gna.org.gh/environment/",
		"https://gna.org.gh/features/",
		"https://gna.org.gh/africa/",
		"https://gna.org.gh/europe/",
		"https://gna.org.gh/americas/",
		"https://gna.org.gh/asias/",
		"https://gna.org.gh/gallery/",
		"https://gna.org.gh/sort/ahafo/",
		"https://gna.org.gh/sort/ashanti/",
		"https://gna.org.gh/sort/bono/",
		"https://gna.org.gh/sort/bono-east/",
		"https://gna.org.gh/sort/central/",
		"https://gna.org.gh/sort/eastern/",
		"https://gna.org.gh/sort/greater-accra/",
		"https://gna.org.gh/sort/northern/",
		"https://gna.org.gh/sort/oti/",
		"https://gna.org.gh/sort/savannah/",
		"https://gna.org.gh/sort/upper-east/",
		"https://gna.org.gh/sort/upper-west/",
		"https://gna.org.gh/sort/volta/",
		"https://gna.org.gh/sort/western/",
		"https://gna.org.gh/sort/western-north/",
		"https://gna.org.gh/foreign-features/",
		"https://gna.org.gh/headline/"})

	// 从翻页器获取链接并访问
	w.OnHTML("a.page-numbers", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从index访问新闻
	w.OnHTML(".entry-header>h2.entry-title>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	// 从index访问新闻
	w.OnHTML(".post-content>h3.entry-title>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	//
	// report.title
	w.OnHTML("h1.entry-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// report .content
	w.OnHTML("div.entry-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = ctx.Content + element.Text
	})
	//report.publish time
	w.OnHTML("header>div.entry-meta>div.date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.author
	w.OnHTML("header>.entry-meta>div.author.by-author", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("#main > div > div.entry-content > p:nth-child(1)", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

}
