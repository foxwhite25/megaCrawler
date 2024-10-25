package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1464", "众议院", "http://www.shugiin.go.jp/")

	engine.SetStartingURLs([]string{
		"https://www.shugiin.go.jp/internet/itdb_annai.nsf/html/statics/topics/topicslist.html",
	})

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

	engine.OnHTML(".topi.mid > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)

		time.Sleep(1 * time.Second) //设置延迟阻断
	})

	engine.OnHTML(".link > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})

	//设置Image的绝对URL
	engine.OnHTML(".pic > a > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		imageUrl := element.Request.AbsoluteURL(element.Attr("src"))
		ctx.Image = []string{imageUrl}
	})

	//采集PDF
	engine.OnHTML(".topien > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
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

	engine.OnHTML(".area > p, .area > table, .right > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
