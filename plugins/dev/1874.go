package dev

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1874", "MAL Contends (Wisconsin)", "https://malcontends.blogspot.com/")

	engine.SetStartingURLs([]string{"https://malcontends.blogspot.com/sitemap.xml"})

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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		switch {
		case strings.Contains(element.Text, "page"):
			engine.Visit(element.Text, crawlers.Index)
		case strings.Contains(element.Text, "20"):
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".post-body > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
