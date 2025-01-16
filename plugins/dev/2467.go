package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2467", "Aero-News Network", "https://www.aero-news.net/")

	engine.SetStartingURLs([]string{"https://www.aero-news.net/emailarticle.cfm?do=main.archive&year=2025"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("div.list-item-news > h2 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)
	})

	engine.OnHTML("#archive-years > li > a, #archive-months > li > a, #archive-days > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})

	engine.OnHTML("div.post > p.date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML("#teaser > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
