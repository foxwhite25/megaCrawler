package storage

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1426", "海峡时报", "https://www.straitstimes.com/global")

	engine.SetStartingURLs([]string{ // 各个新闻板块网站
		"https://www.straitstimes.com/singapore",
		"https://www.straitstimes.com/asia",
		"https://www.straitstimes.com/world",
		"https://www.straitstimes.com/opinion",
		"https://www.straitstimes.com/life",
		"https://www.straitstimes.com/business",
		"https://www.straitstimes.com/tech",
		"https://www.straitstimes.com/sport",
		"https://www.straitstimes.com/st-podcasts",
		"https://www.straitstimes.com/multimedia",
	})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	// 进入各个新闻板块的不同新闻页
	engine.OnHTML(".view-footer > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".card > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	// 处理分页
	engine.OnHTML(".js-pager__items.pagination > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		href := element.Attr("href")
		if strings.Contains(href, "?page=") {
			absoluteURL := element.Request.AbsoluteURL(href)
			engine.Visit(absoluteURL, crawlers.Index)
		}
	})

	// 获取作者信息
	engine.OnHTML(".group-info > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	engine.OnHTML(".layout__region.layout__region--content > div > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
