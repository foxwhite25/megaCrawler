package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2825", "Social Security System", "https://www.sss.gov.ph")

	engine.SetTimeout(60 * time.Second)

	engine.SetStartingURLs([]string{"https://www.sss.gov.ph/news-and-updates/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("#content > div.three_fourth > h1 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("#content > div.three_fourth > div.navigation > ol > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("#content > div.three_fourth > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML("#content > div.three_fourth > div > i", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})
}
