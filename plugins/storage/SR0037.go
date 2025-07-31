package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("SR0037", "Security Solutions Media", "https://www.securitysolutionsmedia.com/")

	engine.SetStartingURLs([]string{
		"https://www.securitysolutionsmedia.com/category/electronic_security/",
		"https://www.securitysolutionsmedia.com/category/physical_security/",
		"https://www.securitysolutionsmedia.com/category/security_management/",
		"https://www.securitysolutionsmedia.com/category/security_operations/",
		"https://www.securitysolutionsmedia.com/category/industry-news/",
	})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("h3 > a,.td-module-thumb > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".fn > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	engine.OnHTML("div.td-post-content.tagdiv-type > p,div.td-post-content.tagdiv-type > ul", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML("td-post-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML("div.page-nav.td-pb-padding-side > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
