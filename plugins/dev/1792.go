package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1792", "哥伦打洛省", "https://gorontaloprov.go.id/")

	engine.SetStartingURLs([]string{"https://berita.gorontaloprov.go.id/wp-sitemap.xml"})

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
		switch {
		case strings.Contains(ctx.URL, "wp-sitemap.xml"):
			engine.Visit(element.Text, crawlers.Index)
		case strings.Contains(ctx.URL, "wp-sitemap-posts"):
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".elementor-container > div:nth-child(1) > div > div:nth-child(3) > div", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
