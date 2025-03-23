package production

import (
	"strings"
	"time"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2406", "国家投资协调署（正部级）", "https://www.bkpm.go.id/")

	engine.SetStartingURLs([]string{"https://www.bkpm.go.id/id/info/siaran-pers"})

	engine.SetTimeout(60 * time.Second) // 延长等待时间

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

	engine.OnHTML(".flex-row.px-1.py-2.row.list-group-item.d-flex.align-items-center > a",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			engine.Visit(element.Attr("href"), crawlers.News)
		})

	engine.OnHTML("li.page-item.border > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("p.pb-1.text-dark-gray.fw-semibold", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(".gap-4.px-4.d-flex.flex-column > div > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
