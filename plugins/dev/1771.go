package dev

import (
	"fmt"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1771", "Pressenza International Press Agency", "https://www.pressenza.com/")

	Sitemap_Part := "https://www.pressenza.com/wp-sitemap-posts-post-"
	Sitemap_Maximum := 10
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
		engine.Visit(element.Text, crawlers.News)
	})

	// engine.OnHTML(".col-xs-12.col-md-8", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	ctx.Content += element.Text
	// })
}
