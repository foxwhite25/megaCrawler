package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1065", "约克郡邮报", "https://www.yorkshirepost.co.uk/")

	engine.SetStartingURLs([]string{
		"https://www.yorkshirepost.co.uk/sitemap.xml",
	})

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
		if strings.Contains(element.Text, "articles") {
			engine.Visit(element.Text, crawlers.Index)
			return
		}
		if strings.Contains(element.Text, ".xml") {
			return
		}
		engine.Visit(element.Text, crawlers.News)
	})

}
