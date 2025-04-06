package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0002", "Cebu Daily News", "https://cebudailynews.inquirer.net/")

	engine.SetStartingURLs([]string{"https://cebudailynews.inquirer.net/category/breaking"})

	engine.SetTimeout(60 * time.Second)

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

	//测试下采集速度较慢
	engine.OnHTML("#pb-info > h1 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("#cdn-m-logo > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("#article-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("noscript").Remove()

		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
