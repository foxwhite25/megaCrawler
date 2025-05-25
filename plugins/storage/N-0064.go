package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0064", "cdc", "https://www.cdc.gov")

	engine.SetStartingURLs([]string{"https://www.cdc.gov/media/site.html"})

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

	engine.OnHTML("ul.list--level-2 > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))

		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)
	})

	engine.OnHTML("span.date-long", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(`meta[property="article:author"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Attr("content"))
	})

	engine.OnHTML("div.dfe-section > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
