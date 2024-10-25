package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1078", "国家利益杂志", "https://nationalinterest.org")

	engine.SetStartingURLs([]string{"https://nationalinterest.org/sitemap.xml"})

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

		case strings.Contains(element.Text, "blog"):
			engine.Visit(element.Text, crawlers.News)

		case strings.Contains(element.Text, "feature"):
			engine.Visit(element.Text, crawlers.News)

		case strings.Contains(element.Text, "commentary"):
			engine.Visit(element.Text, crawlers.News)

		case strings.Contains(element.Text, "article"):
			engine.Visit(element.Text, crawlers.News)

		case strings.Contains(element.Text, "node"):
			engine.Visit(element.Text, crawlers.News)
		}
	})

}
