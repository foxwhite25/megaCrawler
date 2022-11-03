package cna

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"strings"
)

func init() {
	w := megaCrawler.Register("cna", "海军分析中心", "https://www.cna.org/")

	w.SetStartingUrls([]string{"https://www.cna.org/expertise/data-and-architecture-management",
		"https://www.cna.org/expertise/domestic-safety-and-security",
		"https://www.cna.org/expertise/fleet-and-installation-readiness",
		"https://www.cna.org/expertise/force-readiness",
		"https://www.cna.org/expertise/future-fleet-concepts",
		"https://www.cna.org/expertise/global-security-environment",
		"https://www.cna.org/expertise/innovation-and-communications",
		"https://www.cna.org/expertise/plans-and-strategy",
		"https://www.cna.org/expertise/strategic-competition",
		"https://www.cna.org/expertise/tactics-and-operations"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = megaCrawler.Report
		}
	})

	// 从翻页器获取链接并访问
	w.OnHTML("#main-content > section.explore-more-bar > div > div > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})
	w.OnHTML("#main-content > div:nth-child(3) > div > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})

	// 从index访问新闻
	w.OnHTML("#main-content > div > section > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.News)
	})

	//article

	// 添加标题到ctx
	w.OnHTML("#article-header > h1",
		func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
			if ctx.PageType == megaCrawler.Expert {
				ctx.Name = element.Text
			} else if ctx.PageType == megaCrawler.Report || ctx.PageType == megaCrawler.News {
				ctx.Title = element.Text
			}
		})

	// report.author
	w.OnHTML("#article-header > div > span.author", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//report.publish time
	w.OnHTML("#article-header > div > span.dateline", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// 添加正文到ctx
	w.OnHTML("#main-content > div.two-column-layout > article", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})

	//report
	//"https://www.cna.org/expertise/domestic-safety-and-security"
	// 添加标题到ctx
	w.OnHTML("#main-content > div > div.page-title > h1 > a",
		func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
			if ctx.PageType == megaCrawler.Expert {
				ctx.Name = element.Text
			} else if ctx.PageType == megaCrawler.Report || ctx.PageType == megaCrawler.News {
				ctx.Title = element.Text
			}
		})

	// report.author
	w.OnHTML("#main-content > div > div.author-list	", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	// report.description
	w.OnHTML("#main-content > div > section", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Description = element.Text
	})
	//case studies
	// 添加标题到ctx
	w.OnHTML("#main-content > section.above-image-copy > div > h1",
		func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
			if ctx.PageType == megaCrawler.Expert {
				ctx.Name = element.Text
			} else if ctx.PageType == megaCrawler.Report || ctx.PageType == megaCrawler.News {
				ctx.Title = element.Text
			}
		})

	w.OnHTML("#main-content > section.above-image-copy > div > div > p:nth-child(1) > strong", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Description = element.Text
	})

	// 添加前文到ctx
	w.OnHTML("#main-content > section.above-image-copy > div > div", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = ctx.Content + element.Text
	})
	// 添加后文到ctx
	w.OnHTML("//#main-content > section.below-image-copy > div", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = ctx.Content + element.Text
	})

}
