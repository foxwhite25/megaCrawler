package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0060", "name", "https://coltonspointtimes.blogspot.com/")

	engine.SetStartingURLs([]string{
		"https://coltonspointtimes.blogspot.com/sitemap.xml?page=1",
	})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML("div.date-outer > h2 > span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML("div.MsoNormal, div.post-body.entry-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("noscript,script").Remove()
		directText := element.DOM.Text()
		ctx.Content += strings.Join(strings.Fields(directText), " ")
	})
}
