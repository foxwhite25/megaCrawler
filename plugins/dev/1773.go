package dev

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1773", "巴基斯坦官方新闻", "https://en.dailypakistan.com.pk/")

	engine.SetStartingURLs([]string{"https://en.dailypakistan.com.pk/sitemap_index.xml"})

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
		switch {
		case strings.Contains(element.Text, "post-sitemap"):
			engine.Visit(element.Text, crawlers.Index)
		default:
			engine.Visit(element.Text, crawlers.News)
		}
	})

	// engine.OnHTML(".elementor-widget-container", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	ctx.Content += element.Text
	// })
}
