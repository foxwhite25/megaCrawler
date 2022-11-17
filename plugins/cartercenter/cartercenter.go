package cartercenter

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("cartercenter", "卡特中心",
		"https://www.cartercenter.org/")

	w.SetStartingUrls([]string{
		"https://www.cartercenter.org/peace/conflict_resolution/press-releases.html",
		"https://www.cartercenter.org/peace/conflict_resolution/index.html",
		"https://www.cartercenter.org/peace/conflict_resolution/press-releases.html",
		"https://www.cartercenter.org/peace/democracy/press-releases.html",
		"https://www.cartercenter.org/peace/americas/press-releases.html",
		"https://www.cartercenter.org/peace/ati/rule-of-law-press-releases.html",
		"https://www.cartercenter.org/health/index.html",
	})

	// 从 Index 访问 News
	w.OnHTML(".articleTitle>a",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			path := strings.Replace(element.Attr("href"), "../../", "", 1)
			url := "https://www.cartercenter.org/" + path
			w.Visit(url, Crawler.News)
		})

	// 从 /health/index.html 访问 Report
	w.OnHTML("div[class=\"columns four\"]>a[target=\"_self\"]",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			url := "https://www.cartercenter.org/health/" + element.Attr("href")
			w.Visit(url, Crawler.Report)
		})

	// 添加 Title 到 ctx
	w.OnHTML("#brand>h1",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 添加 Content 到 ctx
	w.OnHTML(".wysiwyg",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = strings.TrimSpace(element.Text)
		})

	// 添加 File 到 ctx
	w.OnHTML(".imageWidget>a",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.File = append(ctx.File, element.Attr("href"))
		})

	// 添加 PublicationTime 到 ctx
	w.OnHTML(".articleDate", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
}
