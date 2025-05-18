package dev

import (
	"time"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("ZYT0004", "Mondoweiss", "https://mondoweiss.net")

	engine.SetStartingURLs([]string{"https://mondoweiss.net/news/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.SetTimeout(60 * time.Second)

	engine.OnHTML(".post-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".post-content > :not(div, figure)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".label-next > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
