package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("thz-003", "Dan Mitchell", "https://danieljmitchell.wordpress.com/")

	engine.SetStartingURLs([]string{"https://danieljmitchell.wordpress.com/?s=News"})

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

	engine.OnHTML(".posttitle > h2 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".center > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".entry  p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find(".entry  p > a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
