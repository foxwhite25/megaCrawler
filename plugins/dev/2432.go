package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2432", "AmaBhungane", "https://amabhungane.org/")

	engine.SetTimeout(60 * time.Second) //增加等待时间

	engine.SetStartingURLs([]string{
		"https://amabhungane.org/post-sitemap.xml",
		"https://amabhungane.org/post-sitemap2.xml",
		"https://amabhungane.org/post-sitemap3.xml",
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

	engine.OnHTML("div.elementor-col-66 > div > div > div.elementor-widget-container > p",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
