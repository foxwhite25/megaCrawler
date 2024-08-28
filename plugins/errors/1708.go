package errors

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1708", "国家反贫困委员会", "https://napc.gov.ph/")

	engine.SetStartingURLs([]string{"https://napc.gov.ph/post-sitemap.xml"})

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
		if strings.Contains(ctx.URL, "request-for-quotation") {
		} else {
			engine.Visit(element.Text, crawlers.News)
		}
	})

}
