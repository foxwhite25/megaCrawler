package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0067", "Retail Customer Experience", "https://www.retailcustomerexperience.com/")

	engine.SetStartingURLs([]string{"https://www.retailcustomerexperience.com/news-and-media/"})

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

	engine.OnHTML(".col > h5 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".col-xl.content > p ~ p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".active + a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.Index)
	})
}
