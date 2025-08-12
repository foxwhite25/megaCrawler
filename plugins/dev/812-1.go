
package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
			
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("812-1", "马来西亚广播电视台", "http://www.rtm.gov.my/")
	
	engine.SetStartingURLs([]string{
		"https://www.rtm.gov.my/listings/318",
		"https://www.rtm.gov.my/listings/183",
		"https://www.rtm.gov.my/listings/132",
		"https://www.rtm.gov.my/listings/149"})
	
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
	
	engine.OnHTML(".card-title>a,.table.table-bordered >tbody>tr>td>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("li.active+li", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.Index)
	})

}
