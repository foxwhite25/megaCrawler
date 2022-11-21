package cna

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("cna", "海军分析中心", "https://www.cna.org/")
	w.SetStartingUrls([]string{"https://www.cna.org/experts/default", "https://www.cna.org/our-research/explore-all"})

	//index
	w.OnHTML(" div > div.paginator > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	w.OnHTML("#year-selector > option", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问专家
	w.OnHTML(" div > div > h4 > small > strong > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//获取专家姓名
	w.OnHTML(".col-md-9 > p:nth-child(2) > strong", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	//专家研究领域
	w.OnHTML(" div > div:nth-child(2) > p > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area += element.Text + " "
	})

	//专家描述
	w.OnHTML(" div:nth-child(2) > div > div.col-md-9 > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})

	//专家图片
	w.OnHTML(" div:nth-child(2) > div > div.col-md-3 > div.hidden-xs > img", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Image = []string{element.Attr("href")}
	})

	//访问报告 并获取pdf
	w.OnHTML("#main-content > div > section > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		} else {
			w.Visit(element.Attr("href"), Crawler.Report)
		}
	})

	//报告标题
	w.OnHTML("#main-content > div > div.page-title > h1 > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//作者
	w.OnHTML("#main-content > div > div.author-list > span", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//正文
	w.OnHTML("#main-content > div > section", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Report {
			ctx.Content = element.Text
		}
	})

}
