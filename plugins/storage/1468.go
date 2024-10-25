package storage

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1468", "印度电信管理局", "http://www.trai.gov.in/")

	engine.SetStartingURLs([]string{"https://www.trai.gov.in/whats-new"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	//声明为Report
	engine.OnHTML(".views-field.views-field-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Report)
	})

	engine.OnHTML(".pager-next > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	//采集PDF
	engine.OnHTML(".field-content.press-release-attachement > a, li.first.last > a",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			fileURL := element.Attr("href")
			if strings.Contains(fileURL, ".pdf") {
				url, err := element.Request.URL.Parse(element.Attr("href"))
				if err != nil {
					crawlers.Sugar.Error(err.Error())
					return
				}
				ctx.File = append(ctx.File, url.String())
			}
		})
}
