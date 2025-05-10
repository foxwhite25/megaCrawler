package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("z-0058", "太阳报", "https://www.thesun.co.uk/")

	engine.SetStartingURLs([]string{"https://www.thesun.co.uk/news/uknews/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)
	engine.OnHTML("div.teaser__copy-container > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML("a[rel=\"next\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})
	engine.OnHTML("div.article__content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
