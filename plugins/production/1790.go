package production

import (
	"time"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1790", "西苏门答腊省", "https://www.sumbarprov.go.id/")

	engine.SetStartingURLs([]string{"https://www.sumbarprov.go.id/home/index-berita"})

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
	engine.SetTimeout(60 * time.Second)

	engine.OnHTML(".btn.btn-outline-primary", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("div.card:nth-child(2)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = crawlers.StandardizeSpaces(element.Text)
	})

	engine.OnHTML(".pagination > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.Index)
	})
}
