package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0021", "国土交通省", "http://www.mlit.go.jp/")

	engine.SetStartingURLs([]string{
		"https://www.mlit.go.jp/report/press/houdou202401.html",
		"https://www.mlit.go.jp/report/press/houdou202312.html",
		"https://www.mlit.go.jp/report/press/houdou202212.html",
		"https://www.mlit.go.jp/report/press/houdou202112.html",
		"https://www.mlit.go.jp/report/press/houdou202012.html",
		"https://www.mlit.go.jp/report/press/houdou1912.html",
		"https://www.mlit.go.jp/report/press/houdou1812.html",
		"https://www.mlit.go.jp/report/press/houdou1712.html",
		"https://www.mlit.go.jp/report/press/houdou1612.html",
		"https://www.mlit.go.jp/report/press/houdou1512.html",
		"https://www.mlit.go.jp/report/press/houdou1412.html",
		"https://www.mlit.go.jp/report/press/houdou1312.html",
		"https://www.mlit.go.jp/report/press/houdou1212.html",
		"https://www.mlit.go.jp/report/press/houdou1112.html",
		"https://www.mlit.go.jp/report/press/houdou1012.html",
	})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        false,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".linkPress02 > ul > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("div.text > p > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".titleInner > h2.title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(".clearfix > p.date.mb20", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(".section > p > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
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

	engine.OnHTML("div.section > div.clearfix", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
