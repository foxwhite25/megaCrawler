package dev

import (
	"fmt"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1772", "观察家报", "https://www.spectator.co.uk/")

	// 网站“post-sitemap”太多，用循环建立列表
	Sitemap_Part := "https://www.spectator.co.uk/post-sitemap"
	Sitemap_Maximum := 1
	Sitemaps := []string{}

	for i := 1; i <= Sitemap_Maximum; i++ {
		Sitemap := fmt.Sprintf("%s%d.xml", Sitemap_Part, i)
		Sitemaps = append(Sitemaps, Sitemap)
	}

	engine.SetStartingURLs(Sitemaps)

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
		ctx.Visit(element.Text, crawlers.News)
	})

	// engine.OnHTML(".entry-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	ctx.Content += element.Text
	// })
}
