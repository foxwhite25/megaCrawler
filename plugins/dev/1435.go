package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1435", "马来邮报", "https://www.malaymail.com/")

	engine.SetStartingURLs([]string{
		"https://www.malaymail.com/morearticles/malaysia",
		"https://www.malaymail.com/morearticles/singapore",
		"https://www.malaymail.com/morearticles/money",
		"https://www.malaymail.com/morearticles/world",
		"https://www.malaymail.com/morearticles/life",
		"https://www.malaymail.com/morearticles/tech-gadgets",
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

	engine.OnHTML(".col-md-3.article-item > h2 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("ul.pager > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".article-body > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
