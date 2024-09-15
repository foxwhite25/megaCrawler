package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1030", "东西方中心East_West Center", "https://www.eastwestcenter.org/")

	engine.SetStartingURLs([]string{"https://www.eastwestcenter.org/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)
	// engine.OnResponse((func(response *colly.Response, ctx *crawlers.Context) {
	// 	crawlers.Sugar.Debugln(response.StatusCode)
	// 	crawlers.Sugar.Debugln(string(response.Body))
	// }))
	engine.OnHTML(".link--primary", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
	engine.OnHTML(".pager__item pager__item--next > a ", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
	engine.OnHTML(".teaser__image-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".fullpage__wysiwyg-body  > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	// engine.OnHTML(".pager-next > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	engine.Visit(element.Attr("href"), crawlers.Index)
	// })

	// engine.OnHTML(".pg-normal pg-button pg-next-button", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	engine.Visit(element.Attr("href"), crawlers.Index)
	// })
}
