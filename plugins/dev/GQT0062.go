package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("GQT0062", "Smart", "blog.smart.com.ph")

	engine.SetStartingURLs([]string{
		"https://blog.smart.com.ph/current-events/",
		"https://blog.smart.com.ph/lifestyle/",
		"https://blog.smart.com.ph/pop-culture/",
		"https://blog.smart.com.ph/music/",
		"https://blog.smart.com.ph/sports/",
		"https://blog.smart.com.ph/gaming/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("h2>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".post_header >h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})
	engine.OnHTML(".post_categories+.post_date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
	engine.OnHTML(".post_content>p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	engine.OnHTML(".next", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
	engine.OnHTML(".post_featured with_thumb>img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})
	engine.OnHTML("h2>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Host = element.Text
	})
}
