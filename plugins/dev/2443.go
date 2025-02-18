package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2443", "Bellona", "https://bellona.org/")

	engine.SetTimeout(60 * time.Second)

	engine.SetStartingURLs([]string{
		"https://bellona.org/post-sitemap.xml",
		"https://bellona.org/post-sitemap2.xml",
		"https://bellona.org/post-sitemap3.xml",
		"https://bellona.org/post-sitemap4.xml",
		"https://bellona.org/post-sitemap5.xml",
		"https://bellona.org/post-sitemap6.xml",
		"https://bellona.org/post-sitemap7.xml",
	})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
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

	engine.OnHTML(".gutenberg-simple.wysiwyg.contains-shortcuts > p,.gutenberg-simple.wysiwyg.contains-shortcuts > ul",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
