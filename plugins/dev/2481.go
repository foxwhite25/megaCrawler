package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2481", "Apple Gazette", "https://www.applegazette.com/")

	engine.SetTimeout(60 * time.Second)

	engine.SetStartingURLs([]string{
		"https://www.applegazette.com/post-sitemap.xml",
		"https://www.applegazette.com/post-sitemap2.xml",
		"https://www.applegazette.com/post-sitemap3.xml",
		"https://www.applegazette.com/post-sitemap4.xml",
		"https://www.applegazette.com/post-sitemap5.xml",
		"https://www.applegazette.com/post-sitemap6.xml",
		"https://www.applegazette.com/post-sitemap7.xml",
		"https://www.applegazette.com/post-sitemap8.xml",
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

	engine.OnHTML(`div.elementor-widget-theme-post-content > .elementor-widget-container > p,
		div.elementor-widget-theme-post-content > .elementor-widget-container > ul`,
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			directText := element.DOM.Contents().Not("noscript,a").Text()
			ctx.Content += directText
		})
}
