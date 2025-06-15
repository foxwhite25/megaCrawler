package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("smj-025", "Utah_Democratic_Party", "https://www.utahdemocrats.org/")

	engine.SetStartingURLs([]string{"https://www.utahdemocrats.org/news?offset=1727362800862&reversePaginate=true"})

	extractorConfig := extractors.Config{
		Author:       true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".blog-title>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
		time.Sleep(2 * time.Second)
	})

	engine.OnHTML(".older>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
	engine.OnHTML("div.sqs-html-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
