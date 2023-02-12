package icnl

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("icnl", "非营利组织法国际中心", "https://www.icnl.org/")

	w.SetStartingUrls([]string{"https://www.icnl.org/our-work/freedom-of-association",
		"https://www.icnl.org/our-work/freedom-of-assembly",
		"https://www.icnl.org/our-work/freedom-expression",
		"https://www.icnl.org/our-work/public-participation",
		"https://www.icnl.org/our-work/philanthropy",
		"https://www.icnl.org/coronavirus-response",
		"https://www.icnl.org/our-work/counter-terrorism-security",
		"https://www.icnl.org/our-work/technology-civic-space",
		"https://www.icnl.org/our-work/climate-change-civic-space",
		"https://www.icnl.org/our-work/women-civic-space",
		"https://www.icnl.org/our-work/international-norms-agreements",
		"https://www.icnl.org/our-work/defending-civil-society",
		"https://www.icnl.org/our-work/cross-border-funding",
		"https://www.icnl.org/our-work/civil-society-government-cooperation",
		"https://www.icnl.org/our-work/domestic-fundraising",
		"https://www.icnl.org/our-work/development-civic-space",
		"https://www.icnl.org/our-work/monitoring-assessment",
		"https://www.icnl.org/our-work/asia-pacific-program",
		"https://www.icnl.org/our-work/eurasia-program",
		"https://www.icnl.org/our-work/europe-program",
		"https://www.icnl.org/our-work/latin-america-and-the-caribbean-program",
		"https://www.icnl.org/our-work/mena-program",
		"https://www.icnl.org/our-work/sub-saharan-africa-program",
		"https://www.icnl.org/our-work/us-program",
		"https://www.icnl.org/our-work/global-program",
		"https://www.icnl.org/our-work/civic-space-2040",
		"https://www.icnl.org/coronavirus-response",
		"https://www.icnl.org/news"})

	// 从翻页器获取链接并访问
	w.OnHTML("li.cv-pageitem-number>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})
	w.OnHTML("figure.wpb_wrapper>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})
	// 从index访问新闻
	w.OnHTML("h3.vc_custom_heading>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML("h2.vc_custom_heading>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML(".vc_custom_heading>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML("a.pt-cv-textlink", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML("div.pl-thumbcnt>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	// report.title
	w.OnHTML("main>h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})
	w.OnHTML("h3.news-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})
	// report .content
	w.OnHTML("section.wpb-content-wrapper", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	w.OnHTML("span.read-more-target", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	//report.publish time
	w.OnHTML("h3.vc_custom_heading", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	w.OnHTML("#wrap > h4", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("div.button-download>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})
	w.OnHTML("a.download", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})
	w.OnHTML("div.wpb_wrapper>h6>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})
	// report.author
	w.OnHTML("#wrap > h5", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	// report .content
	w.OnHTML("#wrap", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
}
