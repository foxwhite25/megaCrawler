package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1008", "Philippines News Agency", "https://www.pna.gov.ph/")

	engine.SetStartingURLs([]string{"https://www.pna.gov.ph/"})

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
	engine.OnHTML(".flex > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
	// engine.OnHTML(".just-in-footer > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	engine.Visit(element.Attr("href"), crawlers.Index)
	// })
	engine.OnHTML(".text-2xl > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".article-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	engine.OnHTML(".rounded-e", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
