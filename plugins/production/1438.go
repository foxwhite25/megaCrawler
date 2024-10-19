package production

import (
	"fmt"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1438", "Thailand Ministry of Foreign Affairs", "https://www.mfa.go.th/en")

	baseURL := "https://www.mfa.go.th/en/content-category/press-release"
	totalPages := 344

	// 动态生成所有页面的 URL
	var urls []string
	for page := 1; page <= totalPages; page++ { // 这个网站翻页索引有些不同，导致原本的索引会无效，所以尝试访问所有页面网站。
		url := fmt.Sprintf("%s?p=%d", baseURL, page)
		urls = append(urls, url)
	}

	// 设置所有生成的 URL 作为起始 URL
	engine.SetStartingURLs(urls)

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

	engine.OnHTML("div.jsx-2368373130.content.px-0.py-3.d-flex.flex-column.justify-content-between > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".content-width2.mb-5 > div > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
