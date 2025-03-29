package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2495", "Babel-on-the-Bay", "https://babelonthebay.com/")

	engine.SetStartingURLs([]string{"https://babelonthebay.com/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	//该网站是一个博客，每页包含多篇新闻内容
	engine.OnHTML("ul.page-numbers > li > a.next", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("h2.entry-title,div.entry-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += strings.Join(strings.Fields(element.Text), " ") // 去除换行符、制表符、多余空格
	})
}
