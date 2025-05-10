package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0059", "Colorado Springs Business Journal", "https://www.constructconnect.com/")

	engine.SetStartingURLs([]string{"https://www.constructconnect.com/sitemap.xml"})

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

	engine.OnHTML("div.blog-post__author > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += strings.Join(strings.Fields(element.Text), " ")
	})

	engine.OnHTML("div.blog-post__timestamp", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML("div.blog-post__body > span > p,div.blog-post__body > span > ul", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += strings.Join(strings.Fields(element.Text), " ")
	})
}
