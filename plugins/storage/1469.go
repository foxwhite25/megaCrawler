package storage

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1469", "澳大利亚税务局", "https://www.ato.gov.au")

	engine.SetStartingURLs([]string{
		"https://www.ato.gov.au/sitemap.xml",
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
		if strings.Contains(element.Text, "media-centre") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".AtoContentWrapper_rich-text-content__HY8CB > div > p, .AtoContentWrapper_rich-text-content__HY8CB > div > h2",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
