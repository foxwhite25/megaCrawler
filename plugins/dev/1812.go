package dev

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1812", "Futurity", "https://www.futurity.org/")

	engine.SetStartingURLs([]string{"https://www.futurity.org/sitemap_index.xml"})

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
	// 前几个为xml页面，无新闻，稍等片刻即可出现内容。
	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		switch {
		case strings.Contains(element.Text, "post-sitemap"):
			engine.Visit(element.Text, crawlers.Index)
		default:
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".stickem-container", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
