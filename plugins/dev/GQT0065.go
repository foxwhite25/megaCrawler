package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("GQT0065", "Super Cat", "supercat.ph")

	engine.SetStartingURLs([]string{
		"https://supercat.ph/category/news-and-promos/"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        true,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("h5>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".page-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})
	engine.OnHTML(".post_author>span:nth-child(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	engine.OnHTML(".post_author>span:nth-child(2)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
	engine.OnHTML(".post_main>p,.post_main>ul,h3,h2", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("img").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
	engine.OnHTML(".alignnone,.post_main>img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})
	engine.OnHTML(".breadcrumbs>a:nth-child(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.FirstChannel = element.Text
	})
	engine.OnHTML(".breadcrumbs>a:nth-child(2)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SecondChannel = element.Text
	})
	engine.OnHTML(".breadcrumbs>span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.ThirdChannel = element.Text
	})
	engine.OnHTML(".next.page-numbers", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
