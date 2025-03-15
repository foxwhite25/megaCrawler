package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("3402", "BiggerPockets", "https://www.biggerpockets.com/")

	engine.SetStartingURLs([]string{
		"https://www.biggerpockets.com/blog/post-sitemap1.xml",
		"https://www.biggerpockets.com/blog/post-sitemap2.xml",
		"https://www.biggerpockets.com/blog/post-sitemap3.xml",
		"https://www.biggerpockets.com/blog/post-sitemap4.xml",
		"https://www.biggerpockets.com/blog/post-sitemap5.xml",
		"https://www.biggerpockets.com/blog/post-sitemap6.xml",
		"https://www.biggerpockets.com/blog/post-sitemap7.xml",
		"https://www.biggerpockets.com/blog/post-sitemap8.xml",
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

	engine.OnHTML("#post-content > p, #post-content > h2, #post-content > div > p",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			element.DOM.Find("noscript").Remove()

			directText := element.DOM.Text()
			ctx.Content += directText
		})
}
