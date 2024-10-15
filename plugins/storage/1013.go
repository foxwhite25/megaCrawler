package storage

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1013", "大赦国际", "https://zh.amnesty.org/")

	engine.SetStartingURLs([]string{"https://zh.amnesty.org/"})

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
	engine.OnResponse((func(response *colly.Response, ctx *crawlers.Context) {
		crawlers.Sugar.Debugln(response.StatusCode)
		crawlers.Sugar.Debugln(string(response.Body))
	}))
	engine.OnHTML(".page-nav > ul > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
	engine.OnHTML(".page-numbers > div > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
	engine.OnHTML(".post-figure > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".article-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	// engine.OnHTML(".pager__item > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	//	engine.Visit(element.Attr("href"), crawlers.Index)
	//})
}
