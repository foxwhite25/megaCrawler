package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1299", "简单技术", "https://thetechieguy.com/")

	engine.SetStartingURLs([]string{"https://thetechieguy.com/category/news/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)
	engine.OnHTML(".post_title > h2 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".next", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
	engine.OnHTML("div.the_content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 移除 p 标签中的所有 noscript 标签
		element.DOM.Find("iframe").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
