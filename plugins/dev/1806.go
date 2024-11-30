package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1806", "Electronic Money Association", "https://e-ma.org/")

	engine.SetStartingURLs([]string{"https://e-ma.org/wp-sitemap.xml"})

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
		case strings.Contains(ctx.URL, "wp-sitemap.xml"):
			engine.Visit(element.Text, crawlers.Index)
		case strings.Contains(ctx.URL, "wp-sitemap-posts"):
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".entry-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
