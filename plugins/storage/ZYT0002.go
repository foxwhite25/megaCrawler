package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("ZYT0002", "MobilePaymentsToday", "https://www.paymentsdive.com/")

	engine.SetStartingURLs([]string{"https://www.paymentsdive.com/press-release/"})

	extractorConfig := extractors.Config{
		Author:       true, // 无作者名
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".analytics.t-pr-feed-item", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".content__publish-date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(".content__text", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".pagination > a:nth-last-child(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
