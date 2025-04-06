package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1271z", "Payment Systems Regulator", "https://www.psr.org.uk/")

	engine.SetStartingURLs([]string{"https://www.psr.org.uk/news-and-updates/latest-news/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)
	engine.OnHTML(".m-grid__card > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".m-pagination__controls > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
	engine.OnHTML(".m-rte >p ,.m-rte > ul > li", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	engine.OnHTML("p.m-banner__summary >font,p.m-banner_summary", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += strings.TrimSpace(element.Text)
	})
}
