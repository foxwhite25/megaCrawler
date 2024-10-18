package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1444", "耶路撒冷邮报", "https://www.jpost.com/")

	engine.SetStartingURLs([]string{
		"https://www.jpost.com/jpgooglesitemap/17_SiteMap_Articles_680001-720000.xml",
		"https://www.jpost.com/jpgooglesitemap/18_SiteMap_Articles_720001-760000.xml",
		"https://www.jpost.com/jpgooglesitemap/19_SiteMap_Articles_760001-800000.xml",
		"https://www.jpost.com/jpgooglesitemap/20_SiteMap_Articles_800001-840000.xml",
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

	engine.OnHTML("figure.article-image > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	engine.OnHTML("section.g-row > section > h2", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle += element.Text
	})

	//两种不同风格的article网页
	engine.OnHTML(".article-inner-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(".article-inner-content-breaking-news > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
