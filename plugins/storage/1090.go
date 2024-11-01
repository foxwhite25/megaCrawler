package storage

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1090", "印度联合新闻社", "http://www.uniindia.com/")

	engine.SetStartingURLs([]string{
		"https://www.uniindia.com/news/india/",
		"https://www.uniindia.com/news/world/",
		"https://www.uniindia.com/news/sports/",
		"https://www.uniindia.com/news/business-economy/",
		"https://www.uniindia.com/news/states/",
		"https://www.uniindia.com/news/entertainment/"})

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

	engine.OnHTML(".CatNewsFirst_FirstNews> h1 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".holder > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".storydetails", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = crawlers.StandardizeSpaces(element.Text)
	})
}
