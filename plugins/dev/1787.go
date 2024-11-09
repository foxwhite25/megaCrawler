package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1787", "北苏门答腊省", "https://sumutprov.go.id/")

	engine.SetStartingURLs([]string{"https://sumutprov.go.id/artikel/artikel"})

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

	engine.OnHTML(".blog-item-foot > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".blog-item-body", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(".pagination > li:nth-last-child(2) > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.Index)
	})
}
