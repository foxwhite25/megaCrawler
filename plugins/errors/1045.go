package errors

import (
	"github.com/gocolly/colly/v2"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

// 反爬
func init() {
	engine := crawlers.Register("1045", "莫斯科国立国际关系学院", "https://english.mgimo.ru/news")

	engine.SetStartingURLs([]string{"https://english.mgimo.ru/sitemap.xml"})

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

	engine.OnResponse(func(response *colly.Response, ctx *crawlers.Context) {
		println(string(response.Body))
	})

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Text, crawlers.News)
	})

	// engine.OnHTML(".gutters", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	//	ctx.Content = crawlers.StandardizeSpaces(element.Text)
	//})

	engine.OnHTML(".next > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
