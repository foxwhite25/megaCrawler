package amti

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"strings"
)

func init() {
	w := megaCrawler.Register("amti", "CSIS亚洲海事透明倡议", "https://amti.csis.org/")

	w.SetStartingUrls([]string{"https://amti.csis.org/features/", "https://amti.csis.org/analysis/"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.button", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = megaCrawler.Report
		}
	})

	// 从翻页器获取链接并访问
	w.OnHTML("#main > nav > div > div.nav-previous > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})
	w.OnHTML("#main > nav > div > div.nav-next > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})
	// 从index访问新闻
	w.OnHTML("div > div.col-sm-8 > header > h2 > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.News)
	})

	// report.title
	w.OnHTML("#content > header > div > div > div > div > h1", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})
	//report.publish time
	w.OnHTML(" div.entry-meta", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report .content
	w.OnHTML("div.entry-content", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})
	// report.author
	w.OnHTML(" header > div > a.author.url.fn > font > font", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

}
