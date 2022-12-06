package cnas

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("cnas", "新美国安全中心", "https://www.cnas.org")
	w.SetStartingUrls([]string{"https://www.cnas.org/experts",
		"https://www.cnas.org/articles-multimedia",
		"https://www.cnas.org/reports"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.button.-solid", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})

	//index
	w.OnHTML("section > div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问专家
	w.OnHTML(".h4.margin-vertical-half-em.person-list__name > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//访问文章
	w.OnHTML(".-with-image > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//专家姓名,文章标题
	w.OnHTML("page-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Expert {
			ctx.Name = element.Text
		} else if ctx.PageType == Crawler.Report || ctx.PageType == Crawler.News {
			ctx.Title = element.Text
		}
	})

	//专家头衔
	w.OnHTML("dark medium margin-top-half-em", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//专家描述,文章正文
	w.OnHTML("wysiwyg margin-vertical", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Expert {
			ctx.Description = element.Text
		} else if ctx.PageType == Crawler.Report || ctx.PageType == Crawler.News {
			ctx.Content += element.Text
		}
	})

	//作者
	w.OnHTML("a.contributor", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//时间
	w.OnHTML("sans-serif fz11 bold uppercase margin-bottom-1em\"", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
}
