package opri

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("opri", "海洋政策研究所", "https://www.spf.org/en/opri/")
	w.SetStartingUrls([]string{"https://www.spf.org/en/opri/publication/", "https://www.spf.org/en/opri/news/"})

	//index
	w.OnHTML(" div > div.pagination > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//news
	w.OnHTML(" div.projectYear > div > div > ul > li> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//获取标题
	w.OnHTML("h2.color11", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})
	w.OnHTML("#main > h3", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//获取正文
	w.OnHTML("div > div > div > div > div > div > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content += element.Text
	})

	//report
	w.OnHTML("li.report.clearfix", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		subCtx := ctx.CreateSubContext()
		subCtx.PageType = Crawler.Report
		subCtx.Title = element.ChildText("h4 > a")
		subCtx.File = append(subCtx.File, "h4 > a")
	})
}
