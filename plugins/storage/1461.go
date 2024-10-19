package storage

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1461", "运输安全委员会", "https://www.mlit.go.jp/jtsb/")

	engine.SetStartingURLs([]string{ //这包含两种不同风格的新闻页面
		"https://www.mlit.go.jp/jtsb/houdou.html", //新闻报道
		"https://www.mlit.go.jp/jtsb/kaiken.html", //报道会和相关资料
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

	//根据不同关键字区分不同页面中的Index与News
	engine.OnHTML("#second > ul:not([class]) > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		if strings.Contains(element.Attr("href"), "houdou") && !strings.Contains(element.Attr("href"), ".pdf") {
			engine.Visit(url.String(), crawlers.Index)
		} else if strings.Contains(element.Attr("href"), "kaiken") && !strings.Contains(element.Attr("href"), ".pdf") {
			engine.Visit(url.String(), crawlers.News)
		}

	})

	engine.OnHTML("#second > table> tbody > tr > td > ul > li > ul > li > a",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			if !strings.Contains(element.Attr("href"), ".pdf") {
				engine.Visit(element.Attr("href"), crawlers.News)
			}
		})

	//采集与资料有关的PDF
	engine.OnHTML("#second > ul > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
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

	//不同网站的文本采集（备注：第一类新闻页是以表格的形式）
	engine.OnHTML(".jiko-detail > tbody > tr", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML("#second > p, #second > h3", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
