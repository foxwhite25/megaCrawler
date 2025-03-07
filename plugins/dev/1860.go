package dev

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1860", "Laist", "https://laist.com/")

	engine.SetStartingURLs([]string{"https://laist.com/sitemap.xml"})

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
		case strings.Contains(element.Text, "sitemap"):
			engine.Visit(element.Text, crawlers.Index)
		case strings.Contains(element.Text, "news"):
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".ArticlePage-mainContent", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
