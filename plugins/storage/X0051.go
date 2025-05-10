package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0051", "Patently-O: Patent Law Blog", "https://patentlyo.com/")

	engine.SetStartingURLs([]string{"https://patentlyo.com/"})

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

	engine.OnHTML(".entry-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	// 无法获取标题时去掉注释
	// engine.OnHTML(".entry-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	ctx.Title = element.Text
	// })

	// 部分网站需要登陆才能查看（体现为采到空内容）
	engine.OnHTML(".attributed-text-segment-list__content, .entry-content > p, .mepr-unauthorized-excerpt > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".nav-previous > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.Index)
	})
}
