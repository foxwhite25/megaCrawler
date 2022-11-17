package cdt

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("cdt", "民主与技术中心", "https://cdt.org/")

	w.SetStartingUrls([]string{
		"https://cdt.org/area-of-focus/cybersecurity-standards/",
		"https://cdt.org/area-of-focus/elections-democracy/",
		"https://cdt.org/area-of-focus/equity-in-civic-tech/",
		"https://cdt.org/area-of-focus/free-expression/",
		"https://cdt.org/area-of-focus/government-surveillance/",
		"https://cdt.org/area-of-focus/open-internet/",
		"https://cdt.org/area-of-focus/privacy-data/",
	})

	// 从子频道入口访问 "Read More"
	w.OnHTML(".call-to-action",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Index)
		})

	// 从翻页器访问下一页 Index
	w.OnHTML("a.next-link",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Index)
		})

	// 从 Index 访问报告
	w.OnHTML(".post-archive-item>a",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Report)
		})

	// 添加 Title 到 ctx
	w.OnHTML(".the-title",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 添加 Author 到 ctx.Authors
	w.OnHTML(".the-byline>a",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Authors = append(ctx.Tags, element.Text)
		})

	// 添加 Content 到 ctx
	w.OnHTML("div[class=\"the-content wysiwyg\"]",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = strings.TrimSpace(element.Text)
		})

	// 添加 PublicationTime 到 ctx
	w.OnHTML(".the-date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Attr("datetime")
	})
}
