package dev

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"
)

func init() {
	engine := crawlers.Register("1012", "美国企业研究所", "https://www.aei.org")

	engine.SetStartingURLs([]string{"https://www.aei.org/sitemap_index.xml"})

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
		if strings.Contains(element.Text, "post-sitemap") {
			engine.Visit(element.Text, crawlers.Index)
			return
		}
		if strings.Contains(element.Text, "sitemap") {
			return
		}
		engine.Visit(element.Text, crawlers.News)
	})
}
