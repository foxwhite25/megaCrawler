package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1208", "公正取引委员会", "http://www.jftc.go.jp/")

	engine.SetStartingURLs([]string{"https://www.jftc.go.jp/houdou/pressrelease/index.html"})

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

	engine.OnHTML("body > div.wrapper > div > div.content > ul > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), "index.html") {
			ctx.Visit(element.Attr("href"), crawlers.Index)
		} else {
			ctx.Visit(element.Attr("href"), crawlers.News)
		}
	})

	engine.OnHTML(".content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
}
