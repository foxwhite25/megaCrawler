package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1076", "环球时报", "https://www.globaltimes.cn/")

	engine.SetStartingURLs([]string{
		"https://www.globaltimes.cn/china/index.html",
		"https://www.globaltimes.cn/source/index.html",
		"https://www.globaltimes.cn/opinion/index.html",
		"https://www.globaltimes.cn/world/index.html",
		"https://www.globaltimes.cn/In-depth/index.html",
		"https://www.globaltimes.cn/life/index.html",
		"https://www.dailymail.co.uk/sport/index.html"})

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

	engine.OnHTML(".new_title_ms", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".article_right", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

}
