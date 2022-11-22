package amnesty

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("amnesty", "国际特赦组织", "https://www.amnesty.org/en/")
	w.SetStartingUrls([]string{"https://www.amnesty.org/en/impact/",
		"https://www.amnesty.org/en/research/",
		"https://www.amnesty.org/en/education/",
		"https://www.amnesty.org/en/campaigns/",
		"https://www.amnesty.org/en/latest/news/"})

	//index
	w.OnHTML("a.page-numbers", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//新闻
	w.OnHTML("a.floating-anchor", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//尝试获取pdf并将网页转为报告
	w.OnHTML("a.btn.btn--download.btn--primary", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.PageType = Crawler.Report
			ctx.File = append(ctx.File, element.Attr("href"))
		}
	})

	//获取报告标题
	w.OnHTML("h1.article-attachmentTitle", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//获取新闻标题
	w.OnHTML("h1.article-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//报告时间
	w.OnHTML(".article-attachmentMeta>time", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//新闻时间
	w.OnHTML(".u-textRight", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//正文
	w.OnHTML(" div > section > article > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content += element.Text
	})

	//标签
	w.OnHTML("div.topics-container > ul > li", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})
}
