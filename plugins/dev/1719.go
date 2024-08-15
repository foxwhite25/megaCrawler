package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1719", "农业部", "https://www.da.gov.ph/")

	engine.SetStartingURLs([]string{
		"https://www.da.gov.ph/post-sitemap.xml",
		"https://www.da.gov.ph/post-sitemap2.xml",
		"https://www.da.gov.ph/post-sitemap3.xml",
	})

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
		engine.Visit(element.Text, crawlers.News)
	})

}
