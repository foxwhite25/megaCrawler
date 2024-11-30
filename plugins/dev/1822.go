package dev

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1822", "International Crops Research Institute for the Semi-Arid Tropics", "https://pressroom.icrisat.org/")

	engine.SetStartingURLs([]string{"https://pressroom.icrisat.org/sitemap.xml"})

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
		if strings.Contains(element.Text, "-") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".ContentRenderer_renderer__tPJbs > section", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
