package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0060", "Patent Law Practice Center", "https://patentlawcenter.pli.edu/")

	engine.SetStartingURLs([]string{"https://patentlawcenter.pli.edu/wp-sitemap.xml"})

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
		if strings.Contains(element.Text, "posts-post") {
			engine.Visit(element.Text, crawlers.Index)
		} else if strings.Contains(element.Text, "/20") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".entry > p:not(:last-child)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
