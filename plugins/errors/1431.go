package errors

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1431", "老挝时报", "https://www.vientianetimes.org.la/")

	engine.SetStartingURLs([]string{
		"https://www.vientianetimes.org.la/",
	})

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

	// Error：需要登录
	engine.OnHTML(".section-title > h6 > strong > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("p > strong > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("table > tbody > tr > td > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if element.Attr("class") == "" { // 只处理没有 class 的 <p> 元素
			ctx.Content += element.Text
		}
	})
}
