package errors

import (
	"strings"

	"github.com/gocolly/colly/v2"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

// 反爬
func init() {
	engine := crawlers.Register("1025", "卡托研究所", "https://www.cato.org/sitemaps/default/sitemap.xml")

	engine.SetStartingURLs([]string{"https://www.cato.org/sitemaps/default/sitemap.xml"})

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
		if strings.Contains(element.Text, "sitemap") {
			engine.Visit(element.Text, crawlers.Index)
			return
		}

		engine.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML(".paragraph", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = crawlers.StandardizeSpaces(element.Text)
	})
}
