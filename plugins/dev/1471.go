package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1471", "澳大利亚基础设施与地区发展部", "https://www.infrastructure.gov.au/")

	engine.SetStartingURLs([]string{
		"https://www.infrastructure.gov.au/sitemap.xml?page=1",
		"https://www.infrastructure.gov.au/sitemap.xml?page=2",
		"https://www.infrastructure.gov.au/sitemap.xml?page=3",
		"https://www.infrastructure.gov.au/sitemap.xml?page=4",
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

	//这里NEWS一般是纯新闻，其他是媒体中心的一些报告，通常包含PDF
	keywords := []string{"/media/", "/news/", "/media-centre/"}
	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		for _, keyword := range keywords {
			if strings.Contains(element.Text, keyword) {
				engine.Visit(element.Text, crawlers.News)
				return
			}
		}
	})

	//采集PDF
	engine.OnHTML(".filesize > a, .file > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
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

	engine.OnHTML(".my-5 > p, .my-5 > ul, .pb-5 > p, .node__content.clearfix > div > p,.node__content.clearfix > div > ul",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})

}
