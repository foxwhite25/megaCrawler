package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("z-0079", "INSIDE", "https://www.tsu.edu.ph/")

	engine.SetStartingURLs([]string{"https://www.tsu.edu.ph/news/2025-news/",
		"https://www.tsu.edu.ph/news/2024-news/",
		"https://www.tsu.edu.ph/news/2023-news/",
		"https://www.tsu.edu.ph/news/2022-news/",
		"https://www.tsu.edu.ph/news/2021-news/",
		"https://www.tsu.edu.ph/news/2020-news/",
		"https://www.tsu.edu.ph/news/2019-news/"})

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
	engine.OnHTML("div.has-equal-height >div > a.has-text-dark", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML("a.is-info +a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})
	engine.OnHTML("div.amplify-article > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	engine.OnHTML("h4.has-text-grey", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
}
