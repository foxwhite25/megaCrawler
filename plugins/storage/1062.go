package storage

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"
)

func init() {
	engine := crawlers.Register("1062", "radiotimes.com", "https://www.radiotimes.com/")

	engine.SetStartingURLs([]string{"https://www.radiotimes.com/sitemap.xml"})

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
		if strings.Contains(element.Text, ".xml") {
			engine.Visit(element.Text, crawlers.Index)
			return
		}

		engine.Visit(element.Text, crawlers.News)
	})
}
