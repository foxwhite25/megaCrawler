package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1476", "澳大利亚外交贸易部", "https://www.dfat.gov.au/")

	engine.SetStartingURLs([]string{
		"https://www.dfat.gov.au/sitemap.xml",
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

	keywords := []string{"/news/", "/media-release/", "/international-relations/", "/trade/agreements/"}
	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, ".xml") {
			engine.Visit(element.Text, crawlers.Index)
		} else {
			for _, keyword := range keywords {
				if strings.Contains(element.Text, keyword) {
					engine.Visit(element.Text, crawlers.News)
					return
				}
			}
		}
	})

	// 获取 PublicationTime
	engine.OnHTML(".field > time, .block > span > time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(`.block-region-content > div > div > p,.block-region-content > div > div > ul,
		.block-region-content > div > div > blockquote,
		.paragraph-content > div > p,.paragraph-content > div > h3,
		.contentarea,
		.paragraph-content > div > p`,
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
