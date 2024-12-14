package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2447", "2 Political Junkies (Pennsylvania)", "https://2politicaljunkies.blogspot.com/")

	engine.SetStartingURLs([]string{
		"https://2politicaljunkies.blogspot.com/sitemap.xml?page=1",
		"https://2politicaljunkies.blogspot.com/sitemap.xml?page=2",
		"https://2politicaljunkies.blogspot.com/sitemap.xml?page=3",
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

	engine.OnHTML("h2.date-header > span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(".post-body.entry-content > p, .post-body.entry-content > blockquote, .post-body.entry-content > ul",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
