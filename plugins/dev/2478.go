package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2478", "Android Headlines", "https://www.androidheadlines.com/")

	engine.SetTimeout(60 * time.Second)

	engine.SetStartingURLs([]string{
		"https://www.androidheadlines.com/post-sitemap27.xml",
		"https://www.androidheadlines.com/post-sitemap26.xml",
		"https://www.androidheadlines.com/post-sitemap25.xml",
		"https://www.androidheadlines.com/post-sitemap24.xml",
		"https://www.androidheadlines.com/post-sitemap23.xml",
		"https://www.androidheadlines.com/post-sitemap22.xml",
		"https://www.androidheadlines.com/post-sitemap21.xml",
		"https://www.androidheadlines.com/post-sitemap20.xml",
	})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        false,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Text, crawlers.News)
	})

	// engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
	// 	if strings.Contains(element.Text, "/post-sitemap") {
	// 		engine.Visit(element.Text, crawlers.Index)
	// 	} else if !strings.Contains(element.Text, ".xml") {
	// 		engine.Visit(element.Text, crawlers.News)
	// 	}
	// })

	engine.OnHTML("h1.entry-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML("div.col-12 > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	engine.OnHTML(".entry-content.px-1 > p, .entry-content.px-1 > h2",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
