package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("mwq-006", "Kankakee Daily Journal", "https://www.shawlocal.com/")

	engine.SetStartingURLs([]string{
		"https://www.shawlocal.com/arc/outboundfeeds/sitemap/?outputType=xml",
		"https://www.shawlocal.com/arc/outboundfeeds/sitemap/?outputType=xml&from=100",
		"https://www.shawlocal.com/arc/outboundfeeds/sitemap/?outputType=xml&from=200",
	})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
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

	engine.OnHTML(" div > article > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(" div > section > span.ts-byline__names > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
}
