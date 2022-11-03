package iarc

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"strings"
)

func init() {
	w := megaCrawler.Register("iarc", "北极研究所", "https://uaf-iarc.org/news/")

	w.SetStartingUrls([]string{"https://uaf-iarc.org/news/"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.button", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = megaCrawler.Report
		}
	})
	// report.title
	w.OnHTML("header > h1", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})
	// report .content
	w.OnHTML("div.row.fl-post-image-beside-wrap > div.fl-post-content-beside > div", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})

	// 从翻页器获取链接并访问
	w.OnHTML(" div.fl-builder-pagination-load-more > div > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})
}
