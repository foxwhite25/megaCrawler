package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("ZYT0072", "sjcs", "https://www.sjcs.edu.ph")

	engine.SetStartingURLs([]string{"https://www.sjcs.edu.ph/category/news-and-events/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".link-btn>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("a.next.page-numbers", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".x126k92a>div", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(".date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})

	engine.OnHTML(".breadcrumb>li:nth-child(1)>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.FirstChannel += element.Text
	})

	engine.OnHTML(".breadcrumb>li:nth-child(2)>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SecondChannel += element.Text
	})

	engine.OnHTML(".breadcrumb>li:nth-child(3)>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.ThirdChannel += element.Text
	})

	engine.OnHTML(".cats>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	engine.OnHTML(".author-name", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Tags, strings.TrimSpace(element.Text))
	})
}
