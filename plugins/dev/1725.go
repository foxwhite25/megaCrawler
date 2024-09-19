package dev

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

func init() {
	engine := crawlers.Register("1725", "南华早报", "https://www.scmp.com/")

	engine.SetStartingURLs([]string{"https://www.scmp.com/sitemap.xml"})

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
		if engine.VisitIfContains(element.Text, []string{"article", "hong_kong_25", "economy", "business", "opinion", "news", "tech"}, crawlers.News) {
			return
		}
		engine.Visit(element.Text, crawlers.Index)
	})
}
