package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0009", "The Conversation", "https://theconversation.com")

	engine.SetStartingURLs([]string{"https://theconversation.com/sitemap.xml"})

	engine.SetParallelism(2) //限制进程

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
		if strings.Contains(element.Text, "/au/sitemap_archive") {
			engine.Visit(element.Text, crawlers.Index)
		} else if !strings.Contains(element.Text, ".xml") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("div.content-body.content.entry-content > p,div.content-body.content.entry-content > h2",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
