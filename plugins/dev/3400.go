package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("3400", "Baseline Mag", "https://www.baselinemag.com/")

	engine.SetStartingURLs([]string{
		"https://www.baselinemag.com/post-sitemap.xml",
		"https://www.baselinemag.com/post-sitemap2.xml",
		"https://www.baselinemag.com/post-sitemap3.xml",
		"https://www.baselinemag.com/post-sitemap4.xml",
		"https://www.baselinemag.com/post-sitemap5.xml",
		"https://www.baselinemag.com/post-sitemap6.xml",
		"https://www.baselinemag.com/post-sitemap7.xml",
		"https://www.baselinemag.com/post-sitemap8.xml",
		"https://www.baselinemag.com/post-sitemap9.xml",
		"https://www.baselinemag.com/post-sitemap10.xml",
		"https://www.baselinemag.com/post-sitemap11.xml",
		"https://www.baselinemag.com/post-sitemap12.xml",
	})

	extractorConfig := extractors.Config{
		Author:       true,
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

	engine.OnHTML("div.elementor-widget-container > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	engine.OnHTML("div.elementor-widget-container > p, div.elementor-widget-container > li", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
