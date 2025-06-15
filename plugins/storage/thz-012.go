package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("thz-012", "David Healy Blog", "https://davidhealy.org/")

	engine.SetStartingURLs([]string{"https://davidhealy.org/blog/"})

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

	engine.OnHTML(".entry-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".archive-pagination.pagination li+li >a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".entry-content >p+p,em,strong", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find(".entry-content >p+p>a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

}
