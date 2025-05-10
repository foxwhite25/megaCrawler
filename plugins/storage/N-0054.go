package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0054", "Media Permata", "https://mediapermata.com.bn/")

	engine.SetStartingURLs([]string{"https://mediapermata.com.bn/?s=+"})

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

	engine.OnHTML("#tdi_73 > div > div.td-module-container > div.td-module-meta-info > h3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("div.page-nav > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("div.tdb-category.td-fix-index > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	engine.OnHTML("div.tdb-block-inner.td-fix-index", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += strings.Join(strings.Fields(element.Text), " ")
	})
}
