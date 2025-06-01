package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0070", "Cleaning & Maintenance Management", "https://cmmonline.com/")

	engine.SetStartingURLs([]string{"https://cmmonline.com/news"})

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

	engine.OnHTML("h3.article-block__title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("div.pagination > a.pag-next", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		indexurl := "https://cmmonline.com/news?pag=" + element.Attr("data-page")
		engine.Visit(indexurl, crawlers.Index)
	})

	engine.OnHTML(`meta[property="article:modified_time"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Attr("content"))
	})

	engine.OnHTML("div.article-detail__body > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
