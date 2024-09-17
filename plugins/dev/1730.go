package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1730", "伊朗新闻电视台", "https://www.presstv.ir/")

	engine.SetStartingURLs([]string{"https://www.presstv.ir/Section/13013/"})

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

	engine.OnHTML(".section-latest-news-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	// engine.OnHTML(".col-md-9", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	ctx.Content += element.Text
	// })
	// 若采集到空文章，请将上述三行代码取消注释，并将Text的true改为false。

	engine.OnHTML(".next > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.Index)
	})
}
