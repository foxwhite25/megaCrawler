package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("thz-007", "CryptoCurrencyWire", "https://www.cryptocurrencywire.com/")

	engine.SetStartingURLs([]string{"https://www.cryptocurrencywire.com/cryptonewsbreaks/",
		"https://www.cryptocurrencywire.com/newsroom/newsroom-articles/"})

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
	engine.OnHTML(".newsblurb  a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".nav-links > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".entry-content> p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find(".entry-content > p > a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

}
