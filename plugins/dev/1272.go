package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1272", "PCS", "https://www.pcs.org.uk/")

	engine.SetStartingURLs([]string{"https://www.pcs.org.uk/news-events/news", "https://www.pcs.org.uk/news-events/blogs"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)
	engine.OnHTML(".links > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML("pager__item--next", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href")) //补全为完整URL

		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return //出现错误后打印错误并返回
		}
		engine.Visit(url.String(), crawlers.Index)
	})
	engine.OnHTML(".field--name-name > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text) //补全作者名称
	})
}
