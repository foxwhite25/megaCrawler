package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("SR0006", "PLANSPONSOR", "https://www.plansponsor.com/")

	engine.SetStartingURLs([]string{
		"https://www.plansponsor.com/news/administration/",
		"https://www.plansponsor.com/news/benefits/",
		"https://www.plansponsor.com/news/compliance/",
		"https://www.plansponsor.com/news/deals-and-people/",
		"https://www.plansponsor.com/news/data-and-research/",
		"https://www.plansponsor.com/news/investing/",
		"https://www.plansponsor.com/news/participants/",
		"https://www.plansponsor.com/news/products/"})

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

	engine.OnHTML("h3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("script[type=\"text/javascript\"] + div .align-self-start.article-info:not(span, time)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML("div.content-container.col > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML("a.next-button", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
