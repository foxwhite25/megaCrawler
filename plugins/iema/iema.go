package iema

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("iema", "环境管理与评估研究所", "https://www.iema.net/")
	w.SetStartingUrls([]string{"https://www.iema.net/transform/articles", "https://www.iema.net/resources/news"})

	//index
	w.OnHTML("ul.pagination > li > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})
	w.OnHTML("#topicsNavigation > div > ul > li> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问新闻
	w.OnHTML("a.more", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//新闻标题
	w.OnHTML("div.col-sm-12 > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//新闻正文
	w.OnHTML("div.col-sm-8.mb-6", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	//访问报告
	w.OnHTML("a.btn.btn-primary.blue", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	//报告标题
	w.OnHTML("h1.mt-4.mb-8", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//报告时间
	w.OnHTML(" div.col-md-3 > p > date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//报告分类
	w.OnHTML(".col-md-5.col-lg-push-7.col-lg-offset-1.col-lg-4 > ul> li > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText += element.Text + ","
	})

	//报告标签
	w.OnHTML(".col-lg-push-7.col-lg-offset-1.col-lg-4 > ul.list-unstyled.list-inline > li", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	//报告作者
	w.OnHTML(".col-md-5.col-lg-push-7.col-lg-offset-1.col-lg-4 > p > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//报告正文
	w.OnHTML("div.col-md-pull-6.col-md-6.col-lg-pull-5.col-lg-7.article-body > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content += element.Text
	})
}
