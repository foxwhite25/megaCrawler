package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1293", "世界黄金协会", "https://www.gold.org/")

	engine.SetStartingURLs([]string{"https://www.gold.org/news"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)
	engine.OnHTML(".m-grid > wgc-card > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
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
	engine.OnHTML("div.wgc-text > p , .field-main-content > div > p ", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	engine.OnHTML("time.datetime", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})
}
