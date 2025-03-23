package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2448", "33 Charts", "https://33charts.com/")

	engine.SetStartingURLs([]string{
		"https://33charts.com/post-sitemap.xml",
		"https://33charts.com/post-sitemap2.xml",
	})

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
		engine.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML("div.entry-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 移除 p 标签中的所有 noscript 标签
		element.DOM.Find("noscript").Remove()

		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
