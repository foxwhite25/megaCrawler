package cna

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"strings"
)

func init() {
	w := megaCrawler.Register("cna", "海军分析中心", "https://www.cna.org/")
	w.SetStartingUrls([]string{"https://www.cna.org/experts/default", "https://www.cna.org/our-research/explore-all"})

	//index
	w.OnHTML(" div > div.paginator > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})

	w.OnHTML("#year-selector > option", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})

	//访问专家
	w.OnHTML(" div > div > h4 > small > strong > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Expert)
	})

	//获取专家姓名
	w.OnHTML(".col-md-9 > p:nth-child(2) > strong", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Name = element.Text
	})

	//专家研究领域
	w.OnHTML(" div > div:nth-child(2) > p > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Area += element.Text + " "
	})

	//专家描述
	w.OnHTML(" div:nth-child(2) > div > div.col-md-9 > p", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Description = element.Text
	})

	//专家图片
	w.OnHTML(" div:nth-child(2) > div > div.col-md-3 > div.hidden-xs > img", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Image = []string{element.Attr("href")}
	})

	//访问报告 并获取pdf
	w.OnHTML("#main-content > div > section > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = megaCrawler.Report
		} else {
			w.Visit(element.Attr("href"), megaCrawler.Report)
		}
	})

	//报告标题
	w.OnHTML("#main-content > div > div.page-title > h1 > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})

	//作者
	w.OnHTML("#main-content > div > div.author-list > span", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//正文
	w.OnHTML("#main-content > div > section", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.PageType == megaCrawler.Report {
			ctx.Content = element.Text
		}
	})

}
