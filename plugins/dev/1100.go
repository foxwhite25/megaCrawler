package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1100", "伦敦旗帜晚报", "https://www.timescolonist.com")

	engine.SetStartingURLs([]string{
		"https://www.timescolonist.com/local-news",
		"https://www.timescolonist.com/world-news",
		"https://www.timescolonist.com/2024-bc-votes",
		"https://www.timescolonist.com/national-news",
		"https://www.timescolonist.com/bc-news",
		"https://www.timescolonist.com/animal-stories",
		"https://www.timescolonist.com/national-business",
		"https://www.timescolonist.com/national-sports"})

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

	engine.OnHTML("#category-items > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".selected + li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
