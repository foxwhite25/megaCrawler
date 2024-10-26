package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1073", "DVB", "https://burmese.dvb.no/")

	engine.SetStartingURLs([]string{"https://burmese.dvb.no/sitemap.xml"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "th",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if !strings.Contains(element.Text, "/post/") {
			return
		}
		ctx.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML(".full_content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = crawlers.StandardizeSpaces(element.Text)
	})
}
