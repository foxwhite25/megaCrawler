package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2446", "change.org", "https://www.change.org/")

	engine.SetStartingURLs([]string{"https://www.change.org/sitemap.xml"})

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
		if strings.Contains(element.Text, "/sitemap-20") {
			engine.Visit(element.Text, crawlers.Index)
		} else {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("div.e19irtt30.corgi-1qn3huw > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
