package fpri

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("fpri", "外交政策研究中心", "https://www.fpri.org/")
	w.SetStartingUrls([]string{"https://www.fpri.org/about/scholars/", "https://www.fpri.org/topic/"})

	//访问专家
	w.OnHTML("body > div.rwd-container > div > main > div > section.box-white.about > div > div > div.row > div > ul > li > figcaption > h4 > a",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Expert)
		})

	//获取姓名
	w.OnHTML("div > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Expert {
			ctx.Name = element.Text
		}
	})

	//专家领域
	w.OnHTML(" div.col-md-5 > div > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area = element.Text
	})

	//专家介绍
	w.OnHTML(".inner.dynamic-content > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description += element.Text
	})
	w.OnHTML(".inner.dynamic-content > div > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description += element.Text
	})

	//访问新闻
	w.OnHTML("span.btn.btn-blue", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	w.OnHTML("a.next.page-numbers", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	w.OnHTML("a.readmore", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//获取新闻标题
	w.OnHTML("h2.caption-article", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//获取作者
	w.OnHTML("li > a.author", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//获取时间
	w.OnHTML("li>span", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.News {
			ctx.PublicationTime = element.Text
		}
	})

	//分类
	w.OnHTML("div.list-cat", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = element.Text
	})

	//正文
	w.OnHTML("body > div.rwd-container > div > main > div > section > div > div > div.row > div:nth-child(1) > p",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content += element.Text
		})
}
