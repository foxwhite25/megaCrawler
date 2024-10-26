package storage

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1463", "原子能规制委员会", "https://www.nra.go.jp/")

	engine.SetStartingURLs([]string{
		"https://www.da.nra.go.jp/search?ftxt=1&fuse=1&f.t=10",
		"https://www.da.nra.go.jp/search?ftxt=1&fuse=1&f.t=11",
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

	engine.OnHTML(".c-list-item-pc__labels > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)
	})

	engine.OnHTML(".c-pagination__count-wrapper > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})

	// 根据PDF的链接构建新的URL，构建后的URL为直接下载该PDF
	engine.OnHTML(".p-details-mix__item-content > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fileURL := element.Attr("href")
		if strings.Contains(fileURL, "view") {
			absURL, err := element.Request.URL.Parse(fileURL)
			if err != nil {
				crawlers.Sugar.Error(err.Error())
				return
			}

			// 提取 ?contents= 之后的部分作为文件名
			queryIdx := strings.Index(absURL.String(), "?contents=")
			if queryIdx != -1 {
				// 提取 ?contents= 之后的内容
				fileName := absURL.String()[queryIdx+10:]

				// 构建新的PDF下载链接
				downloadURL := "https://www.da.nra.go.jp/data/" + fileName + ".pdf"

				url, err := element.Request.URL.Parse(downloadURL)
				if err != nil {
					crawlers.Sugar.Error(err.Error())
					return
				}

				ctx.File = append(ctx.File, url.String())
				ctx.PageType = crawlers.Report
			}
		}
	})

	engine.OnHTML(".p-details-event__head > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
