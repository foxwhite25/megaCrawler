package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0070", "Ridenbaugh Press/publishing", "https://www.ridenbaugh.com/")

	engine.SetStartingURLs([]string{"https://www.ridenbaugh.com/"})

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

	engine.OnHTML(".post-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".post-byline", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(".post-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".next.page-numbers", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.Index)
	})
}
