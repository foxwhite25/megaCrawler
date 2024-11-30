package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2429", "African Liberty", "https://www.africanliberty.org/")

	engine.SetStartingURLs([]string{
		"https://www.africanliberty.org/post-sitemap.xml",
		"https://www.africanliberty.org/post-sitemap2.xml",
		"https://www.africanliberty.org/post-sitemap3.xml",
		"https://www.africanliberty.org/post-sitemap4.xml",
		"https://www.africanliberty.org/post-sitemap5.xml",
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

	engine.OnHTML(`div.elementor-element.elementor-element-62ca394.elementor-widget.elementor-widget-theme-post-content > div > p,
	div.elementor-element.elementor-element-62ca394.elementor-widget.elementor-widget-theme-post-content > div > h2`,
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			// 过滤noscript标签内容,避免可能的垃圾信息
			directText := element.DOM.Contents().Not("noscript").Text()
			ctx.Content += directText
		})
}
