package jiia

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("jiia", "日本国际问题研究所", "https://www.jiia.or.jp/en/")
	w.SetStartingUrls([]string{"https://www.jiia.or.jp/en/tags/", "https://www.jiia.or.jp/en/abus/experts.html"})

	//index
	w.OnHTML("#main > section.section.tags-list-wrap > div > ul > li > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})
	w.OnHTML("page-link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问报告
	w.OnHTML("#fs-result > section > div > ul > li> article > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	//获取标题
	w.OnHTML("#main > section > div > div.title-box > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//获取时间
	w.OnHTML("#main > section > div > div.title-box > dl > dt", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//获取作者
	w.OnHTML("#main > section > div > div.title-box > dl > dd", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//获取正文
	w.OnHTML("#main > section > div > div.post-contents > div > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	//获取分类
	w.OnHTML("div.cat-name", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = element.Text
	})

	//pdf
	w.OnHTML("div.link-to-pdf", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})

	//访问专家
	w.OnHTML(" strong > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//name
	w.OnHTML("#main > section > div > div > div > div:nth-child(1) > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	//头衔
	w.OnHTML("#main > section > div > div > div > div:nth-child(1) > div", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//描述
	w.OnHTML("#main > section > div > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})
}
