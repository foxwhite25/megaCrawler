package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1772", "观察家报", "https://www.spectator.co.uk/")

	engine.SetStartingURLs([]string{"https://www.spectator.co.uk/sitemap_index.xml"})

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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Request.URL.String(), "sitemap_index") {
			if strings.Contains(element.Text, "post") {
				ctx.Visit(element.Text, crawlers.Index)
				return
			}
			return
		}
		ctx.Visit(element.Text, crawlers.News)
	})
}
