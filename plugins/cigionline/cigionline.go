package cigionline

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("cigionline", "国际治理创新中心", "https://www.cigionline.org/")

	w.SetStartingUrls([]string{"https://www.cigionline.org/research/",
		"https://www.cigionline.org/topics/africa/", "https://www.cigionline.org/topics/artificial-intelligence/",
		"https://www.cigionline.org/topics/big-data/", "https://www.cigionline.org/topics/china/",
		"https://www.cigionline.org/topics/competition/", "https://www.cigionline.org/topics/democracy/",
		"https://www.cigionline.org/topics/digital-currency/", "https://www.cigionline.org/topics/emerging-technology/",
		"https://www.cigionline.org/topics/financial-systems/", "https://www.cigionline.org/topics/future-work/",
		"https://www.cigionline.org/topics/g20g7/", "https://www.cigionline.org/topics/gender/",
		"https://www.cigionline.org/topics/india/", "https://www.cigionline.org/topics/innovation/",
		"https://www.cigionline.org/topics/innovation-economy/", "https://www.cigionline.org/topics/intellectual-property/",
		"https://www.cigionline.org/topics/internet-governance/", "https://www.cigionline.org/topics/platform-governance/",
		"https://www.cigionline.org/topics/security/", "https://www.cigionline.org/topics/space/",
		"https://www.cigionline.org/topics/standards/", "https://www.cigionline.org/topics/surveillance-privacy/",
		"https://www.cigionline.org/topics/systemic-risk/", "https://www.cigionline.org/topics/trade/",
		"https://www.cigionline.org/opinion-series/", "https://www.cigionline.org/opinions/",
		"https://www.cigionline.org/events/"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.button", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})

	//翻页是由按钮来操控的，不懂
	// 从index访问新闻
	w.OnHTML("a.table-title-link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	// report.title

	w.OnHTML("div.col > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//report.publish time
	w.OnHTML("section > div > div > div > div.date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	w.OnHTML("div.col > div:nth-child(3)", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report .content
	w.OnHTML("section.body", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	//内含Expert
	w.OnHTML("a.block-author", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})
	// expert.Name
	w.OnHTML(" div.col-md-8.col-lg-6.hero-main > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})
	// expert.description
	w.OnHTML("div.short-bio>p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = ctx.Content + element.Text
	})
	// expert.description
	w.OnHTML("#article-body", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = ctx.Content + element.Text
	})
	// expert.area
	w.OnHTML("span.expertise-item", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area = ctx.Area + "," + element.Text
	})

	// expert.link
	w.OnHTML("div.contact>div", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})

}
