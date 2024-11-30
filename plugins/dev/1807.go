package dev

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"
)

func init() {
	engine := crawlers.Register("1807", "Equity", "https://www.equity.org.uk/")

	engine.SetStartingURLs([]string{"https://www.equity.org.uk/sitemap.xml"})

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
		if !strings.Contains(element.Text, "/news/") {
			return
		}
		engine.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML(".c-wysiwyg", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
