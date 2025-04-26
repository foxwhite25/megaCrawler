package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0038", "Invest Slovenia.", "https://www.sloveniabusiness.eu/")

	engine.SetStartingURLs([]string{"https://www.sloveniabusiness.eu/sitemap.xml"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "hot-topics") || (strings.Contains(element.Text, "business-news")) {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	// engine.OnHTML(".fw-bold.cfs20", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	ctx.Title = element.Text
	// })

	engine.OnHTML(".container > time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(".cfs10 p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
