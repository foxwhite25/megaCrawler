package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("tribune_1", "The Express Tribune", "https://tribune.com.pk/")

	engine.SetStartingURLs([]string{"https://tribune.com.pk/sitemap.xml"})

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
		if strings.Contains(element.Text, "/sitemap/posts") {
			engine.Visit(element.Text, crawlers.Index)
		} else if !strings.Contains(element.Text, ".xml") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("div.left-authorbox > span:nth-child(2)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML("div.storypage-rightside > span.story-text", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("noscript,script").Remove()
		directText := element.DOM.Text()
		ctx.Content += strings.Join(strings.Fields(directText), " ")
	})
}
