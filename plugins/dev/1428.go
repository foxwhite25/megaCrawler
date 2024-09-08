package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1428", "曼谷邮报", "http://www.bangkokpost.com/")

	engine.SetStartingURLs([]string{"https://www.bangkokpost.com/sitemap_news.xml"})

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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Text, crawlers.News)
	})

	//采集网站的副标题（如果它有的话）
	engine.OnHTML(".article-headline > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle += element.Text
	})

	engine.OnHTML(".article-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
