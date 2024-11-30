package dev

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2422", "Times of Oman", "https://m.timesofoman.com/")

	engine.SetStartingURLs([]string{
		"https://m.timesofoman.com/category/oman",
		"https://m.timesofoman.com/category/world",
		"https://m.timesofoman.com/category/business",
		"https://m.timesofoman.com/category/sports",
		"https://m.timesofoman.com/category/opinion",
		"https://m.timesofoman.com/category/technology",
		"https://timesofoman.com/category/roundup",
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

	engine.OnHTML("h2.post-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(`li.page-item > a[rel="next"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("span.text-muted", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML("#__content__ > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
