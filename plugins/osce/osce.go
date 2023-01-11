package osce

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("osce", "欧洲安全与合作组织", "https://www.osce.org/")

	w.SetStartingUrls([]string{"https://www.osce.org/stories",
		"https://www.osce.org/resources/publications",
		"https://www.osce.org/press-releases",
		"https://www.osce.org/countering-terrorism",
		"https://www.osce.org/conflict-prevention-and-resolution",
		"https://www.osce.org/cyber-ict-security",
		"https://www.osce.org/democratization",
		"https://www.osce.org/economic-activities",
		"https://www.osce.org/education",
		"https://www.osce.org/elections",
		"https://www.osce.org/environmental-activities",
		"https://www.osce.org/gender-equality",
		"https://www.osce.org/good-governance",
		"https://www.osce.org/human-rights",
		"https://www.osce.org/media-freedom-and-development",
		"https://www.osce.org/migration",
		"https://www.osce.org/national-minority-issues",
		"https://www.osce.org/policing",
		"https://www.osce.org/reform-and-cooperation-in-the-security-sector",
		"https://www.osce.org/roma-and-sinti",
		"https://www.osce.org/rule-of-law",
		"https://www.osce.org/tolerance-and-nondiscrimination",
		"https://www.osce.org/youth",
		"https://www.osce.org/sustainable-development-goals",
		"https://www.osce.org/resources/csce-osce-key-documents",
		"https://www.osce.org/resources/documents/decision-making-bodies",
		"https://www.osce.org/resources/documents",
		"https://www.osce.org/resources/publications",
		"https://www.osce.org/press-releases",
		"https://www.osce.org/resources/multimedia",
		"https://www.osce.org/resources/e-libraries",
		"https://www.osce.org/resources/126378",
		"https://www.osce.org/osce-social-media",
		"https://www.osce.org/cca",
		"https://www.osce.org/resources/documents"})

	// 从翻页器获取链接并访问
	w.OnHTML("li.pager-item>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})
	w.OnHTML("li.pager-next>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})
	w.OnHTML("li.pager-last>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})
	w.OnHTML("div.link-item-link>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从index访问新闻
	w.OnHTML(".osce-figure-content>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML("dt.search-title>h3>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})
	w.OnHTML("div.pane-content>div.field-field-item-source>div>div>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})
	w.OnHTML("span.field-content>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	// report.title
	w.OnHTML("div.content--header>h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})
	w.OnHTML("div.content--content>h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})
	//report.publish time
	w.OnHTML("span.date-display-single", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.author
	w.OnHTML("span.author", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// report .content
	w.OnHTML("div.content--content>div>div.pane-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	w.OnHTML("div.node-content>div.pane-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	// report.description
	w.OnHTML("div.field-field-description>div>div.field-item", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})
	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("ul.osce-download-links>li>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})
}
