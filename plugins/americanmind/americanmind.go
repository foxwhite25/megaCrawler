package americanmind

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("americanmind", "克萊蒙研究所", "https://americanmind.org/")

	w.SetStartingUrls([]string{"https://americanmind.org/salvos/",
		"https://americanmind.org/memos/",
		"https://americanmind.org/features/",
		"https://americanmind.org/audio/",
		"https://americanmind.org/video/"})

	// 从翻页器获取链接并访问
	w.OnHTML("", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从index访问新闻
	w.OnHTML("h1.tam__post-content-title-output>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML("a.article-linkb", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	// report.title
	w.OnHTML("div.tam__single-header-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})
	//report.publish time
	w.OnHTML("span.tam__single-header-meta-type", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	// report.author
	w.OnHTML("div.tam__single-header-author>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	// report .content
	w.OnHTML("div.tam__single-content-area", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
}
