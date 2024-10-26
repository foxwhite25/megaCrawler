package storage

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1462", "海上保安厅", "http://www.kaiho.mlit.go.jp/")

	engine.SetStartingURLs([]string{"https://www.kaiho.mlit.go.jp/info/kouhou/"})

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

	engine.OnHTML(".titleBasic > h3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	// 采集PDF
	engine.OnHTML(".text_container__inner.rich-text > p > a, font > a, span > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fileURL := element.Attr("href")
		if strings.Contains(fileURL, ".pdf") {
			url, err := element.Request.URL.Parse(element.Attr("href"))
			if err != nil {
				crawlers.Sugar.Error(err.Error())
				return
			}
			ctx.File = append(ctx.File, url.String())
			ctx.PageType = crawlers.Report
		}
	})

	engine.OnHTML(".text_container__inner.rich-text > p, .text_container__inner.rich-text > h3", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
