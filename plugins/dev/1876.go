package dev

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1876", "MarketersMEDIA", "https://mm-newsroom-459236083887.us-central1.run.app/")

	engine.SetStartingURLs([]string{"https://mm-newsroom-459236083887.us-central1.run.app/sitemap.xml"})

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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "sitemap") {
			engine.Visit(element.Text, crawlers.Index)
		} else if !strings.Contains(element.Text, "sitemap") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".f18.single-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
