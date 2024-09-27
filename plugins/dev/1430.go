package dev

import (
	"fmt"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"regexp"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1430", "新加坡商业时报", "https://www.businesstimes.com.sg/")

	// 创建一个切片来存储所有的sitemap URL
	startingURLs := []string{}

	// 使用循环来生成多个sitemap的URL
	for i := 1; i <= 84; i++ {
		url := fmt.Sprintf("https://www.businesstimes.com.sg/_plat/api/v1/sitemap.xml?page=%d", i)
		startingURLs = append(startingURLs, url) // 将生成的 URL 添加到切片
	}

	// 设置起始URL
	engine.SetStartingURLs(startingURLs)

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

	// 这是匹配新闻的正则表达式
	newsURLPattern := regexp.MustCompile(`^https://www\.businesstimes\.com\.sg/[a-zA-Z0-9/-]+$`)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		url := element.Text
		// 只在满足表达式的情况下进行匹配
		if newsURLPattern.MatchString(url) {
			engine.Visit(url, crawlers.News)
		}
	})

	engine.OnHTML(".container.px-0.mx-auto.w-full.mb-4 > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle += element.Text
	})

	engine.OnHTML(".relative.mb-4 > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
