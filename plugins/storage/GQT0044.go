package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("GQT0044", "QCT Online", "https://www.qctonline.com/")

	engine.SetStartingURLs([]string{
		"https://www.qctonline.com/category/arts-entertainment/",
		"https://www.qctonline.com/category/features/",
		"https://www.qctonline.com/category/city-news/",
		"https://www.qctonline.com/category/community/",
		"https://www.qctonline.com/category/the-province/",
		"https://www.qctonline.com/category/opinion/",
		"https://www.qctonline.com/category/street-views/",
		"https://www.qctonline.com/category/environment/",
	})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        true,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("h1 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".author.vcard > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	engine.OnHTML(".tags-links > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	engine.OnHTML(".entry-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".next", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
