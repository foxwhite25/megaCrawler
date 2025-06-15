package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("the-006", "Crypto Breaking News", "https://www.cryptobreaking.com/")

	engine.SetStartingURLs([]string{"https://www.cryptobreaking.com/category/news/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".is-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".main-pagination > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".post-content.cf.entry-content.content-spacious > p,.post-content.cf.entry-content.content-spacious  > h2", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find(".post-content.cf.entry-content.content-spacious > p > a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
