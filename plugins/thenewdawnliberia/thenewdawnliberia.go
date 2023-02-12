package thenewdawnliberia

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("thenewdawnliberia", "新黎明", "https://thenewdawnliberia.com/")

	w.SetStartingUrls([]string{"https://thenewdawnliberia.com/category/oped/",
		"https://thenewdawnliberia.com/category/ecowas-news/",
		"https://thenewdawnliberia.com/category/other-news/health/",
		"https://thenewdawnliberia.com/category/features/crime-a-punishment/",
		"https://thenewdawnliberia.com/category/liberia-political-hotfire-with-jones-mallay/",
		"https://thenewdawnliberia.com/category/editorial/",
		"https://thenewdawnliberia.com/category/features/commentary-features/",
		"https://thenewdawnliberia.com/category/other-news/rural-news/",
		"https://thenewdawnliberia.com/category/business/finance-business/",
		"https://thenewdawnliberia.com/category/features/opinion-features/",
		"https://thenewdawnliberia.com/category/liberia-news/environmental-news/",
		"https://thenewdawnliberia.com/category/liberia-news/politics-news/",
		"https://thenewdawnliberia.com/category/features/special-feature-features/",
		"https://thenewdawnliberia.com/category/liberia-news/francais/",
		"https://thenewdawnliberia.com/category/business/investment-business/",
		"https://thenewdawnliberia.com/category/liberia-news/ama/",
		"https://thenewdawnliberia.com/africa-news/"})

	// 从翻页器获取链接并访问
	w.OnHTML(".last-page>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})
	// 从index访问新闻
	w.OnHTML("a.more-link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	// report.title
	w.OnHTML("h1.post-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// report .content
	w.OnHTML("div.entry-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = ctx.Content + element.Text
	})
	//report.publish time
	w.OnHTML(".entry-header>.post-meta>.date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.author
	w.OnHTML("#the-post > div.entry-content.entry.clearfix > p:nth-child(2)", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

}
