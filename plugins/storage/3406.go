package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("3406", "Bladen Journal (Elizabethtown, North Carolina)", "https://www.bladenjournal.com/")

	engine.SetStartingURLs([]string{"https://www.bladenjournal.com/sitemap.xml"})

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
		if strings.Contains(element.Text, "news") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("div.td-post-featured-image img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	engine.OnHTML("#content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
