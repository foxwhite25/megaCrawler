package storage

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	// 国土交通省下属的外局
	engine := crawlers.Register("1459", "观光厅", "https://www.mlit.go.jp/kankocho/")

	engine.SetStartingURLs([]string{
		// 这个网站存留了2021-2024年的内容，2020年以前的内容被留在了旧网站中
		"https://www.mlit.go.jp/kankocho/news_2024.html",
		"https://www.mlit.go.jp/kankocho/news_2023.html",
		"https://www.mlit.go.jp/kankocho/news_2022.html",
		"https://www.mlit.go.jp/kankocho/news_2021.html",
		"https://www.mlit.go.jp/kankocho/topics_2024.html",
		"https://www.mlit.go.jp/kankocho/topics_2023.html",
		"https://www.mlit.go.jp/kankocho/topics_2022.html",
		"https://www.mlit.go.jp/kankocho/topics_2021.html",
	})

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

	engine.OnHTML(".js-pullDownFilterContents > div > ul > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	// 获取时间
	engine.OnHTML(".st-article__cont > p.update", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 采集PDF
	engine.OnHTML(".st-article__elm > ul.c-list > li > a, .st-article__head > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
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

	engine.OnHTML("div.st-article__cont > div.st-article__cont", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
