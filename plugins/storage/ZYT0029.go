package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("ZYT0029", "dailynexus", "https://dailynexus.com")

	engine.SetStartingURLs([]string{"https://dailynexus.com/category/news"})

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

	engine.OnHTML(".category-page-post-text>h1>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".page-navigation>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(strings.ToLower(element.Text), "next") {
			engine.Visit(element.Attr("href"), crawlers.Index)
		}
	})

	engine.OnHTML(".single-post-content>p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

}
