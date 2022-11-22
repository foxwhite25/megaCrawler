package pemsea

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("pemsea", "东亚海洋环境治理伙伴关系", "http://www.pemsea.org/")
	w.SetStartingUrls([]string{"http://pemsea.org/publications/reports", "http://pemsea.org/news"})

	//index
	w.OnHTML(" div.item-list > ul > li> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//news
	w.OnHTML(" div.views-field.views-field-body > div > div > div:nth-child(1) > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//report
	w.OnHTML("div > div.col-sm-10 > div.title > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	//标题
	w.OnHTML("page-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})
	//时间
	w.OnHTML("p.submitted", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//正文
	w.OnHTML(" div > div > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content += element.Text
	})

	//pdf
	w.OnHTML("a.pub-link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})
}
