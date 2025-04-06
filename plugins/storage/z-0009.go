package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("z-0009", "The Huddersfield Daily Examiner", "https://www.examinerlive.co.uk/")

	engine.SetStartingURLs([]string{"https://www.examinerlive.co.uk/news/"})

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
	engine.OnHTML(".teaser-text > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML("div.article-body >p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
