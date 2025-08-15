package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("ZYT0050", "STVPlayer", "https://news.stv.tv")

	engine.SetStartingURLs([]string{
		"https://news.stv.tv/section/scotland",
		"https://news.stv.tv/section/west-central",
		"https://news.stv.tv/section/east-central",
		"https://news.stv.tv/section/north",
		"https://news.stv.tv/section/highlands-islands",
		"https://news.stv.tv/section/politics",
		"https://news.stv.tv/section/entertainment",
		"https://news.stv.tv/section/world",
	})

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

	engine.OnHTML(".frontpage-section .headline-container > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("#main > .subhead", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle += element.Text
	})

	engine.OnHTML(".article-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(".next", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
