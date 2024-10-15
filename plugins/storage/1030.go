package storage

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"
)

func init() {
	engine := crawlers.Register("1030", "东西方中心East_West Center", "https://www.eastwestcenter.org/")

	engine.SetStartingURLs([]string{"https://www.eastwestcenter.org/sitemap.xml"})

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
		if strings.Contains(element.Text, "/publication/") {
			engine.Visit(element.Text, crawlers.News)
		}
		if strings.Contains(element.Text, "sitemap") {
			engine.Visit(element.Text, crawlers.Index)
		}
	})
}
