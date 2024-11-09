package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1200", "亚太安全研究中心", "https://edmontonjournal.com")

	engine.SetStartingURLs([]string{
		"https://edmontonjournal.com/category/news/local-news/",
		"https://edmontonjournal.com/category/news/politics/",
		"https://edmontonjournal.com/category/health/",
		"https://edmontonjournal.com/category/news/crime/",
		"https://edmontonjournal.com/category/news/true-crime/",
		"https://edmontonjournal.com/category/news/national/",
		"https://edmontonjournal.com/category/news/world/"})

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

	engine.OnHTML(".article-card__link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

}
