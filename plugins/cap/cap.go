package cap

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"strings"
)

func init() {
	w := megaCrawler.Register("cna", "美国进步中心", "https://www.americanprogress.org/")
	w.SetStartingUrls([]string{"https://www.americanprogress.org/issues/#topics", "https://www.americanprogress.org/experts/"})

	//访问人物
	w.OnHTML(" div.people1 > div.-o\\:h > div > article> a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Expert)
	})

	//人物名称
	w.OnHTML(" div > header > div.header1-main > h1", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.PageType == megaCrawler.Expert {
			ctx.Name = element.Text
		}
	})

	//领域
	w.OnHTML(" div > header > div.header1-main > h3", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Area += element.Text + " "
	})

	//专家介绍
	w.OnHTML(" div > div.-xw\\:4.-mx\\:a.-mb\\:3 > div > p", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Description = element.Text
	})

	//专家图片
	w.OnHTML("div.header1-side > figure > span.img1.-top > img", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Image = []string{element.Attr("href")}
	})

	//index
	w.OnHTML(" div:nth-child(1) > div > div.-o\\:h > div > div", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})
	w.OnHTML("#showmore > div > ul > li> a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})
	w.OnHTML("div.archives1-wrap > p > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})

	//从index访问report
	w.OnHTML(" div.archives1-wrap > div > article > div > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Report)
	})

	//标题
	w.OnHTML(" header > div.header2-wrap > div.header2-main > h1", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})

	//报告摘要
	w.OnHTML(" div > header > div.header2-wrap > div.header2-main > p", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Description = element.Text
	})

	//作者
	w.OnHTML(" div.header2-wrap > div.header2-side > div.authors1 > ul > li", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	//时间
	w.OnHTML(" div > header > div.header2-brow > time", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})
	//类型
	w.OnHTML(" div > header > div.header2-brow > span > span", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.CategoryText = element.Text
	})

	//正文
	w.OnHTML("#content > div > div > div.article1", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})

}
