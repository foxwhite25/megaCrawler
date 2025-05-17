package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("thz-010", "Conwebwatch", "https://conwebwatch.tripod.com")

	engine.SetStartingURLs([]string{"https://conwebwatch.tripod.com/archive.html"})

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

	engine.OnHTML("td:nth-child(1) > h5 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("td:nth-child(1) > p,blockquote", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		cleanText := strings.ReplaceAll(element.Text, "\n", "")
		cleanText = strings.ReplaceAll(cleanText, "\t", "")
		ctx.Content += cleanText
	})

}
