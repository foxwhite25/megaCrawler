package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1418", "越南网", "http://vietnamnet.vn/")

	engine.SetStartingURLs([]string{"http://vietnamnet.vn/"})

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

	// 访问所有新闻页
	engine.OnHTML(".header_submenu-content > ul > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".horizontalPost__main > h3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".pagination__list > li.pagination__list-item > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".content-detail.content-mobile-change > h2", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description += element.Text
	})

	engine.OnHTML(".maincontent.main-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
