package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0024", "BSR", "https://www.bsr.org/")

	engine.SetStartingURLs([]string{"https://www.bsr.org/en/blog"})

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

	engine.OnHTML("div.card-body > h2 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("ul.pagination.justify-content-center.pt-4 > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("div.col > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		imageUrl := element.Request.AbsoluteURL(element.Attr("src"))
		ctx.Image = []string{imageUrl}
	})

	engine.OnHTML(".col-md-7.order-md-2.order-1 > p,.col-md-7.order-md-2.order-1 > ul", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
