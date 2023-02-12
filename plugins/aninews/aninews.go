package aninews

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("aninews", "亚洲国际新闻", "https://aninews.in/")

	w.SetStartingUrls([]string{"https://aninews.in/category/national/general-news/",
		"https://aninews.in/category/national/politics/",
		"https://aninews.in/category/national/features/",
		"https://aninews.in/category/world/asia/",
		"https://aninews.in/category/world/us/",
		"https://aninews.in/category/world/europe/",
		"https://aninews.in/category/world/pacific/",
		"https://aninews.in/category/world/others/",
		"https://aninews.in/category/world/middle-east/",
		"https://aninews.in/category/business/corporate/",
		"https://aninews.in/category/science/",
		"https://aninews.in/category/tech/mobile/",
		"https://aninews.in/category/tech/internet/",
		"https://aninews.in/category/tech/computers/",
		"https://aninews.in/category/tech/others/",
		"https://aninews.in/category/environment/",
		"https://aninews.in/category/lifestyle/culture/",
		"https://aninews.in/category/health/"})

	// 从翻页器获取链接并访问
	w.OnHTML("li.page-elem-wrap>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从index访问新闻
	w.OnHTML(".content>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML("figcaption>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	// report.title
	w.OnHTML("h1[itemprop=\"headline\"]", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// report .content
	w.OnHTML("article>div.content:nth-child(2)", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	//report.publish time
	w.OnHTML("span.time-red[itemprop=\"dateModified\"]", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
}
