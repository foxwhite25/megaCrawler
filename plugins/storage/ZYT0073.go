package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("ZYT0073", "pnb", "https://www.pnb.com.ph")
	engine.SetStartingURLs([]string{"https://www.pnb.com.ph/index.php/news?tpl=2&utm_campaign=News&utm_medium=bitly&utm_source=Website+-+Homepage+Desktop"})
	extractorConfig := extractors.Config{
		Author:       false, //no author
		Image:        false,
		Language:     true,
		PublishDate:  false, //no publish date
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}
	extractorConfig.Apply(engine)

	engine.OnHTML(".page-header>h2>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".content.clearfix>div>p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".imgstory.with-caption>img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})
}
