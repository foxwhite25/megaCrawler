package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1065", "约克郡邮报", "https://www.yorkshirepost.co.uk/")

	engine.SetStartingURLs([]string{
		"https://www.yorkshirepost.co.uk/news/latest",
		"https://www.yorkshirepost.co.uk/news/politics",
		"https://www.yorkshirepost.co.uk/business",
		"https://www.yorkshirepost.co.uk/education",
		"https://www.yorkshirepost.co.uk/health",
		"https://www.yorkshirepost.co.uk/news/transport",
		"https://www.yorkshirepost.co.uk/news/crime",
		"https://www.yorkshirepost.co.uk/news/world",
		"https://www.yorkshirepost.co.uk/news/uk-news",
	})

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

	engine.OnHTML(".article-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".article-content > div > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

}
