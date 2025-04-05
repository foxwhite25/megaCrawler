package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1290", "UKFINANCE.UK", "https://www.ukfinance.org.uk/")

	engine.SetStartingURLs([]string{"https://www.ukfinance.org.uk/news-and-insight"})

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
	engine.OnHTML(".card_title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML("a[title=\"Go to next page\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href")) //补全为完整URL

		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return //出现错误后打印错误并返回
		}
		engine.Visit(url.String(), crawlers.Index)
	})
	engine.OnHTML("div.content > div p,div.content > div>ol >li", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	engine.OnHTML("div.pg-intro__text > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Text
	})
}
