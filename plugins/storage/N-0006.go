package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0006", "Báo Bình Dương", "https://baobinhduong.vn/")

	engine.SetStartingURLs([]string{
		"https://baobinhduong.vn/chinh-tri/",
		"https://baobinhduong.vn/kinh-te/",
		"https://baobinhduong.vn/quoc-te/",
		"https://baobinhduong.vn/xa-hoi/",
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

	engine.OnHTML("div.news-cate-middle > div:nth-child(1) > div.box-latestnews > div.items-news div.content > h3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("div.paginate-items > button.item", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 获取当前请求的URL
		currentURL := element.Request.URL.String()

		urlPath := currentURL

		// 如果URL中包含分页参数，去掉这部分
		if idx := strings.Index(urlPath, "?page="); idx >= 0 {
			urlPath = urlPath[:idx]
		}
		if !strings.HasSuffix(urlPath, "/") {
			urlPath += "/"
		}

		// 组合新的分页URL
		url := urlPath + "?page=" + element.Text
		engine.Visit(url, crawlers.Index)
	})

	engine.OnHTML("div.singular-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
