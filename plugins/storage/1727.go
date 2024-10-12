package storage

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1727", "国际商业时报", "https://www.ibtimes.com/")

	engine.SetStartingURLs([]string{"https://www.ibtimes.com/world"})

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

	engine.OnHTML(".image > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".paper-next > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
