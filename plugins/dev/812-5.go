
package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
		
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("812-5", "马来西亚消防与拯救局", "http://www.bomba.gov.my/")
	
	engine.SetStartingURLs([]string{
		"https://www.bomba.gov.my/category/uncategorized/",
		"https://www.bomba.gov.my/aktiviti/"})
	
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
	
	engine.OnHTML(".entry-title.ast-blog-single-element>a,.elementor-tab-content>ul>li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".entry-content.clear > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(".next.page-numbers", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.Index)
	})
}
