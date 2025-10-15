
package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("smj925-4", "National-Grid-Corporation-of-the-Philippines", "http://www.napocor.gov.ph/")
	
	engine.SetStartingURLs([]string{"https://www.napocor.gov.ph/"})
	
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
	engine.OnHTML(".entry-title>a ", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	
	engine.OnHTML(".nav-previous>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
