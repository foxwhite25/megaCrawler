package nf

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("nf", "日本财团", "https://www.nippon-foundation.or.jp/en")
	w.SetStartingUrls([]string{"https://www.nippon-foundation.or.jp/en/news/articles/"})

	//index
	w.OnHTML("li.Pager__item.Pager__item--next > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})
	w.OnHTML("li.LocalNavItem>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//news
	w.OnHTML("a.NewsItem__inner", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//获取标题
	w.OnHTML(".RTE.SingleH1>h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//时间
	w.OnHTML("time.NewsItem__date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//正文
	w.OnHTML(" div.RTE.RTE--single", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

}
