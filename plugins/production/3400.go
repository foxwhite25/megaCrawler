package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("3400", "Baseline Mag", "https://undergroundreporter.org/")

	engine.SetStartingURLs([]string{
		"https://undergroundreporter.org/post-sitemap.xml",
		"https://undergroundreporter.org/post-sitemap2.xml",
		"https://undergroundreporter.org/post-sitemap3.xml",
		"https://undergroundreporter.org/post-sitemap4.xml",
		"https://undergroundreporter.org/post-sitemap5.xml",
		"https://undergroundreporter.org/post-sitemap6.xml",
		"https://undergroundreporter.org/post-sitemap7.xml",
		"https://undergroundreporter.org/post-sitemap8.xml",
		"https://undergroundreporter.org/post-sitemap9.xml",
		"https://undergroundreporter.org/post-sitemap10.xml",
		"https://undergroundreporter.org/post-sitemap11.xml",
		"https://undergroundreporter.org/post-sitemap12.xml",
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
		// 移除 p 标签中的所有 noscript 标签
		element.DOM.Find("noscript").Remove()

		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
