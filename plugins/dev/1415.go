package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1415", "越南快讯", "https://www.vnexpress.net")

	engine.SetStartingURLs([]string{ //六个板块的新闻页
		"https://vnexpress.net/thoi-su/chinh-tri",
		"https://vnexpress.net/thoi-su/dan-sinh",
		"https://vnexpress.net/thoi-su/lao-dong-viec-lam",
		"https://vnexpress.net/thoi-su/giao-thong",
		"https://vnexpress.net/thoi-su/mekong",
		"https://vnexpress.net/thoi-su/quy-hy-vong",
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

	engine.OnHTML(".item-news.item-news-common > h3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
		time.Sleep(1000 * time.Millisecond) //每次访问延迟1000ms，避免请求太频繁被服务器屏蔽
	})

	engine.OnHTML("a.btn-page", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("div.sidebar-1 > p.description", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description += element.Text
	})

	engine.OnHTML("article.fck_detail > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

}
