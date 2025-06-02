package dev

import (
	"fmt"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("smj-011", "经济发展研究所", "https://www.ide.go.jp/")

	baseURL := "https://www.ide.go.jp/English/ResearchColumns.html"
	totalPages := 5

	// 动态生成所有页面的 URL
	var urls []string
	for page := 1; page <= totalPages; page++ {
		url := fmt.Sprintf("%s?_page=%d", baseURL, page)
		urls = append(urls, url)
	}

	// 设置所有生成的 URL 作为起始 URL
	engine.SetStartingURLs(urls)

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".p-card-01.p-2column-02.record > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".paragraph.pbNested.pbNestedWrapper ", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML("p[style*='text-align: right;']>br+br,p[style*='text-align: right;']", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "2") {
			// 获取元素内容或属性
			ctx.PublicationTime = element.Text
		}
	})

	engine.OnHTML("p[style*='text-align: right;']>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})
}
