package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1026", "民主中心", "https://democracy.uchicago.edu/news/")

	engine.SetStartingURLs([]string{"https://democracy.uchicago.edu/news/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)
	// engine.OnResponse((func(response *colly.Response, ctx *crawlers.Context) {
	// 	crawlers.Sugar.Debugln(response.StatusCode)
	// 	crawlers.Sugar.Debugln(string(response.Body))
	// }))
	engine.OnHTML(".post-content >  a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
	engine.OnHTML(".entry-content > p > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	// engine.OnHTML(".sitemap > tbody > tr > td > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	engine.Visit(element.Attr("href"), crawlers.News)
	// })
	// engine.OnHTML(".post-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	engine.Visit(element.Attr("href"), crawlers.Report)
	// })
	engine.OnHTML(".article__abs > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	// engine.OnHTML(".pg-normal pg-button pg-next-button", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	engine.Visit(element.Attr("href"), crawlers.Index)
	// })
}
