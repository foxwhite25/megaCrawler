package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("z-0078", "INFONT", "https://infotechlead.com/")

	engine.SetStartingURLs([]string{"https://infotechlead.com/?s=+news"})

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
	engine.OnHTML("div.td-module-meta-info > h3.entry-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("span.current + a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})

	engine.OnHTML("div.tdb-block-inner > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 移除 p 标签中的所有 noscript 标签
		element.DOM.Find("img").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText

	})
}
