﻿package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1277", "POAUK.ORG", "https://www.poauk.org.uk/")

	engine.SetStartingURLs([]string{"https://www.poauk.org.uk/news-events/news-room/"})

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
	engine.OnHTML(".listing  > div > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".next", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
	engine.OnHTML("#left-col > p, ul > li  em", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 移除标签中的所有无关标签
		element.DOM.Find("img").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
