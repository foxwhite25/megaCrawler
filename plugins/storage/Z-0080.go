package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("z-0080", "IMO", "https://www.imo.org/")

	engine.SetStartingURLs([]string{
		"https://www.imo.org/zh/mediacentre/pressbriefings/pages/2024-archives.aspx",
		"https://www.imo.org/zh/mediacentre/pressbriefings/pages/2023-archives.aspx",
		"https://www.imo.org/zh/mediacentre/pressbriefings/pages/2022-archives.aspx",
		"https://www.imo.org/zh/mediacentre/pressbriefings/pages/2021-archives.aspx",
		"https://www.imo.org/zh/mediacentre/pressbriefings/pages/2020-archives.aspx",
		"https://www.imo.org/zh/mediacentre/pressbriefings/pages/2019-archives.aspx",
		"https://www.imo.org/zh/mediacentre/pressbriefings/pages/default.aspx"})

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
	engine.OnHTML("div.card-body > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML("li.next> a.page-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})

	engine.OnHTML("div.heading-subtitle", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
}
