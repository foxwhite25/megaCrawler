package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("ZYT0043", "Splash247", "https://splash247.com")

	engine.SetStartingURLs([]string{"https://splash247.com/sitemap.xml"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "post.20") {
			engine.Visit(element.Text, crawlers.Index)
		} else if !strings.Contains(element.Text, "xml") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".entry-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("img").Remove()
		element.DOM.Find("span").Remove()
		element.DOM.Find("script").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
