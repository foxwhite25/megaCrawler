package hoover

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"regexp"
	"strings"
)

func init() {
	w := megaCrawler.Register("hoover", "胡佛研究所", "https://www.hoover.org/")

	w.SetStartingUrls([]string{"https://www.hoover.org/research/topic/monetary-policy",
		"https://www.hoover.org/research/topic/us-labor-market",
		"https://www.hoover.org/research/topic/budget-spending",
		"https://www.hoover.org/research/topic/trade",
		"https://www.hoover.org/research/topic/finance-banking",
		"https://www.hoover.org/topic/k-12",
		"https://www.hoover.org/topic/higher-education",
		"https://www.hoover.org/research/topic/nuclear-energy",
		"https://www.hoover.org/research/topic/renewable-energy",
		"https://www.hoover.org/research/topic/health-care",
		"https://www.hoover.org/research/topic/china",
		"https://www.hoover.org/research/topic/europe",
		"https://www.hoover.org/research/topic/russia",
		"https://www.hoover.org/research/topic/indo-pacific",
		"https://www.hoover.org/research/topic/middle-east",
		"https://www.hoover.org/research/topic/africa",
		"https://www.hoover.org/research/topic/democracy",
		"https://www.hoover.org/research/topic/us-foreign-policy",
		"https://www.hoover.org/research/topic/us-defense",
		"https://www.hoover.org/research/topic/arms-control",
		"https://www.hoover.org/topic/cyber-security",
		"https://www.hoover.org/research/topic/intelligence",
		"https://www.hoover.org/research/topic/terrorism",
		"https://www.hoover.org/research/topic/immigration",
		"https://www.hoover.org/research/topic/innovation",
		"https://www.hoover.org/research/topic/intellectual-property",
		"https://www.hoover.org/research/topic/us",
		"https://www.hoover.org/research/topic/military",
		"https://www.hoover.org/research/topic/world",
		"https://www.hoover.org/research/topic/contemporary",
		"https://www.hoover.org/research/topic/campaigns-elections",
		"https://www.hoover.org/research/topic/political-philosophy",
		"https://www.hoover.org/research/topic/congress",
		"https://www.hoover.org/research/topic/public-opinion",
		"https://www.hoover.org/research/topic/presidency",
		"https://www.hoover.org/research/topic/comparative-politics",
		"https://www.hoover.org/research/topic/judiciary",
		"https://www.hoover.org/research/topic/civil-rights-race",
		"https://www.hoover.org/research/topic/regulation-property-rights",
		"https://www.hoover.org/research/topic/technology-law-and-governance",
		"https://www.hoover.org/research/topic/fiscal-policy"})

	//https://www.hoover.org/research/topic/monetary-policy,https://www.hoover.org/research/topic/us-labor-market
	//通过翻页器去访问所有index
	//翻页
	w.OnHTML("#pagination > div > ul > li.ais-Pagination-item.ais-Pagination-item--nextPage > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})
	//访问report
	w.OnHTML("#hits > div > div > ol > li:nth-child > div > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Report)
	})
	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = megaCrawler.Report
		}
	})

	//report .descption
	//CONGRESSIONAL TESTIMONY
	//workpaper
	w.OnHTML("#block-hoover-content > section.article-detail.news-wrap.research-detail.small-font > div > div > div.content", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Description = element.Text
	})
	//report.CategoryText
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.bg-light-gray > div.container > div > div.col-left > div.content-wrap > span.article > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.CategoryText = element.Text
	})
	//report .title
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.bg-light-gray > div.container > div > div.col-left > div.content-wrap > h1 > p", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})

	//report.publish time
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.bg-light-gray > div.container > div > div.col-left > div.content-wrap > span.date > font:nth-child(1) > font", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//report.author
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.bg-light-gray > div.container > div > div.col-left > div.content-wrap > span.author-info.view > font > a > font", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	//report.author
	w.OnHTML("#block-hoover-content > section.article-detail.news-wrap.research-detail.small-font > div > div > div.sidebar.article-sidebar > div.content-description > p", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//report.description
	w.OnHTML("#block-hoover-content > section.article-detail.news-wrap.research-detail.small-font > div > div > div.content ", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Description = element.Text
	})

	//report.link
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.bg-light-gray > div.container > div > div.col-left > div.content-wrap > ul > li > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})
	//另一个下载按钮
	w.OnHTML(" #block-hoover-content > section.article-detail.news-wrap.research-detail.small-font > div > div > div.content > p > strong > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
		ctx.PageType = megaCrawler.Report
	})

	//Essays
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.bg-light-gray > div.container > div > div.col-left > div.content-wrap > span.article > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		eventsRegex, _ := regexp.Compile("\\w+")
		events := eventsRegex.FindStringSubmatch(element.Text)
		if len(events) == 2 {
			ctx.CategoryText = events[1]
		}
	})
	//Content
	w.OnHTML("#block-hoover-content > section.article-detail.news-wrap.event-detail.small-font", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.CategoryText == "Essays" {
			ctx.Content = element.Text
		}
	})
	//title
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.news-detail.events-detail > div.container > div > div.col-left > div > h1", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.CategoryText == "Essays" {
			ctx.Title = element.Text
		}
	})
	//description
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.news-detail.events-detail > div.container > div > div.col-left > div > div.text-wrap", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.CategoryText == "Essays" {
			ctx.Description = element.Text
		}
	})
	//publish time
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.news-detail.events-detail > div.container > div > div.col-left > div > span.date", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.CategoryText == "Essays" {
			ctx.PublicationTime = element.Text
		}
	})
	//author
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.news-detail.events-detail > div.container > div > div.col-left > div > span.author-info > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.CategoryText == "Essays" {
			ctx.Authors = append(ctx.Authors, element.Text)
		}
	})
	//
	//link
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.news-detail.events-detail > div.container > div > div.col-left > div > div.btn-wrapper > ul > li > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.CategoryText == "Essays" {
			ctx.Link = append(ctx.Link, element.Attr("href"))
		}
	})

	//另一个下载按钮
	w.OnHTML(" div.toolbar_drop > div > div.toolbar_bottom > div.toolbar_items>a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
		ctx.PageType = megaCrawler.Report
	})

	//article
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.bg-light-gray > div.container > div > div.col-left > div.content-wrap > span.article > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.CategoryText = element.Text
	})
	//Content
	w.OnHTML("#block-hoover-content > section.article-detail.news-wrap.research-detail.small-font", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.CategoryText == "Articles" {
			ctx.Content = element.Text
		}
	})

	//title
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.bg-light-gray > div.container > div > div.col-left > div.content-wrap > h1", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.CategoryText == "Articles" {
			ctx.Title = element.Text
		}
	})
	//description
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.bg-light-gray > div.container > div > div.col-left > div.content-wrap > div > p", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.CategoryText == "Articles" {
			ctx.Description = element.Text
		}
	})
	//publish time
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.bg-light-gray > div.container > div > div.col-left > div.content-wrap > span.date", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.CategoryText == "Articles" {
			ctx.PublicationTime = element.Text
		}
	})
	//author
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.bg-light-gray > div.container > div > div.col-left > div.content-wrap > span.author-info.view > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.CategoryText == "Articles" {
			ctx.Authors = append(ctx.Authors, element.Text)
		}
	})
	//
	//link
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.bg-light-gray > div.container > div > div.col-left > div.content-wrap > ul > li > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.CategoryText == "Articles" {
			ctx.Link = append(ctx.Link, element.Attr("href"))
		}
	})

	//events  Events
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.bg-light-gray > div.container > div > div.col-left > div.content-wrap > span.article > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		eventsRegex, _ := regexp.Compile("\\w+")
		events := eventsRegex.FindStringSubmatch(element.Text)
		if len(events) == 2 {
			ctx.CategoryText = events[1]
		}
	})
	//Content
	w.OnHTML("#block-hoover-content > section.article-detail.news-wrap.event-detail.small-font", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.CategoryText == "Events" {
			ctx.Content = element.Text
		}
	})
	//title
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.news-detail.events-detail > div.container > div > div.col-left > div > h1", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.CategoryText == "Events" {
			ctx.Title = element.Text
		}
	})
	//description
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.news-detail.events-detail > div.container > div > div.col-left > div > div.text-wrap", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.CategoryText == "Events" {
			ctx.Description = element.Text
		}
	})
	//publish time
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.news-detail.events-detail > div.container > div > div.col-left > div > span.date", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.CategoryText == "Events" {
			ctx.PublicationTime = element.Text
		}
	})
	//author
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.news-detail.events-detail > div.container > div > div.col-left > div > span.author-info > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.CategoryText == "Events" {
			ctx.Authors = append(ctx.Authors, element.Text)
		}
	})
	//
	//link
	w.OnHTML("#block-hoover-content > section.banner-with-detail.padding-small-top.news-detail.events-detail > div.container > div > div.col-left > div > div.btn-wrapper > ul > li > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.CategoryText == "Events" {
			ctx.Link = append(ctx.Link, element.Attr("href"))
		}
	})

}
