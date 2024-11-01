package storage

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1099", "英国太阳报", "https://www.thesun.co.uk/")

	engine.SetStartingURLs([]string{"https://www.thesun.co.uk/sitemap.xml"})

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
		switch {
		case strings.Contains(element.Text, "sitemap"):
			engine.Visit(element.Text, crawlers.Index)

		default:
			engine.Visit(element.Text, crawlers.News)
		}
	})
}
