package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1729", "马来西亚国家通讯社", "https://www.trtworld.com/")

	engine.SetStartingURLs([]string{"https://trtworld.com/sitemap-articles.xml"})

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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Text, crawlers.News)
	})

	// engine.OnHTML(".Article-List", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	ctx.Content += element.Text
	// })
	// 若采集到空文章，请将上述三行代码取消注释，并将Text的true改为false。
}
