package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("GQT0023", "More Than Shipping", "https://www.morethanshipping.com/")

	engine.SetStartingURLs([]string{
		"https://www.morethanshipping.com/topics/news/",
		"https://www.morethanshipping.com/topics/supply-chain-planning/",
		"https://www.morethanshipping.com/topics/maritime-shipping/",
		"https://www.morethanshipping.com/topics/ports/",
		"https://www.morethanshipping.com/topics/rail/",
		"https://www.morethanshipping.com/topics/trucking/",
		"https://www.morethanshipping.com/topics/business/",
		"https://www.morethanshipping.com/topics/shipping-trends/"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("h3>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".tdb-block-inner.td-fix-index>time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
	engine.OnHTML(".tdb-author-name", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	engine.OnHTML(".tdb-block-inner.td-fix-index>p,.tdb-block-inner.td-fix-index>ul", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	engine.OnHTML(".page-nav.td-pb-padding-side>a:nth-child(6)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
