package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1810", "FICCI", "https://www.ficci.in/api/home")

	engine.SetStartingURLs([]string{"https://www.ficci.in/api/press_releases"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        false,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".bg-orange", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".info > h5", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	engine.OnHTML(".info", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
