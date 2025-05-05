package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0049", "Kampuchea Thmey Daily", "http://www.kampucheathmey.com")

	engine.SetStartingURLs([]string{"https://www.kampucheathmey.com/sitemap_index.xml"})

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
		if strings.Contains(element.Text, "/post-sitemap") {
			engine.Visit(element.Text, crawlers.Index)
		} else if !strings.Contains(element.Text, ".xml") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("span.author-name.author > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	engine.OnHTML("#mvp-content-main > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
