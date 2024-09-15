package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1028", "沃森国际公共事务研究所关于战争经费的研究", "https://watson.brown.edu/costsofwar/news")

	engine.SetStartingURLs([]string{"https://watson.brown.edu/costsofwar/news"})

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
	// engine.OnHTML(".leaf > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	engine.Visit(element.Attr("href"), crawlers.Index)
	// })
	engine.OnHTML(".newsitem  > div > a ", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	// engine.OnHTML(".post-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	engine.Visit(element.Attr("href"), crawlers.Report)
	// })
	engine.OnHTML(".read__content  > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	engine.OnHTML(".pager-next > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	// engine.OnHTML(".pg-normal pg-button pg-next-button", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	engine.Visit(element.Attr("href"), crawlers.Index)
	// })
}
