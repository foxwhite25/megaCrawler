package dev

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1864", "Law360", "https://www.law360.com/")

	engine.SetStartingURLs([]string{"https://www.law360.com/sitemap.xml"})

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
		case strings.Contains(element.Text, "articles_"):
			engine.Visit(element.Text, crawlers.Index)
		case strings.Contains(element.Text, "articles/"):
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".entry-content > div", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
