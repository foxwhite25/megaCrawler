package storage

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1466", "日本银行", "http://www.boj.or.jp/")

	engine.SetStartingURLs([]string{"https://www.boj.or.jp/whatsnew/index.htm"})

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

	engine.OnHTML(".panel-cont > ul > li > a, .outline.mod_outer > p > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".htm") {
			url, err := element.Request.URL.Parse(element.Attr("href"))
			if err != nil {
				crawlers.Sugar.Error(err.Error())
				return
			}
			engine.Visit(url.String(), crawlers.News)
		}
	})

	// 采集PDF（包括索引页的所有PDF文件）
	engine.OnHTML(".panel-cont > ul > li > a, .link-list01 > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
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

	engine.OnHTML(".outline.mod_outer", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
