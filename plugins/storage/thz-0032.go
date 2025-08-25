package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("thz-0032", " ", "https://beta.entrepreneurship.org.ph/")

	engine.SetStartingURLs([]string{"https://beta.entrepreneurship.org.ph/category/staff-news/"})

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

	engine.OnHTML(".entry-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".col-lg-12.col-md-12.col-xs-12.col-sm-12 > p,.row span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find(".col-lg-12.col-md-12.col-xs-12.col-sm-12 > p > img").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".nav-previous > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
