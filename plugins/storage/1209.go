package storage

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1209", "消防厅", "http://www.fdma.go.jp/")

	engine.SetStartingURLs([]string{
		"https://www.fdma.go.jp/pressrelease/houdou/",
		"https://www.fdma.go.jp/pressrelease/info/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        false,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".txt-list a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		subCtx := ctx.CreateSubContext()
		subCtx.File = append(subCtx.File, element.Attr("href"))
		subCtx.Title = element.Text
		subCtx.PageType = crawlers.Report
	})

	engine.OnHTML(".side-list > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.Index)
	})
}
