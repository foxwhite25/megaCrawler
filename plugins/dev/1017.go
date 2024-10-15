package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1017", "卡托研究所", "https://www.cato.org/search/category/event")

	engine.SetStartingURLs([]string{"https://www.cato.org/search/category/event"})

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
	// engine.OnHTML(".sub-menu__item > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	engine.Visit(element.Attr("href"), crawlers.Index)
	// })
	engine.OnHTML(".pager-item > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
	engine.OnHTML(".mb-3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".body-text > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	// 	engine.OnHTML(".card-grid__load-more  > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 		engine.Visit(element.Attr("href"), crawlers.Index)
	// 	})
}
