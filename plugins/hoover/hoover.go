package hoover

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("hoover", "胡佛研究所", "https://www.hoover.org")
	w.SetStartingUrls([]string{"https://www.hoover.org/research/type/essays",
		"https://www.hoover.org/fellows",
		"https://www.hoover.org/research/type/working-papers"})

	//尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("strong", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})

	//访问人物
	w.OnHTML("div > div.hover-content > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//获取人物名字
	w.OnHTML(" section.banner-with-video > div.container > div.text-wrap > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	//专家介绍
	w.OnHTML(" div > div.content-wrap > div.text-wrapper > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})

	//专家领域
	w.OnHTML(" div > div.sidebar-wrap > div > div:nth-child(1) > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Text, "Expertise") {
			w.OnHTML(" div.sidebar-wrap > div > div:nth-child(1) > ul > li", func(element *colly.HTMLElement, ctx *Crawler.Context) {
				ctx.Area += element.Text + " "
			})
		}
	})

	//index
	w.OnHTML("#pagination > div > ul > li > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//从index访问报告
	w.OnHTML("#hits > div > div > ol > li > div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	//报告类型
	w.OnHTML(" div.col-left > div.content-wrap > span.article > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = element.Text
	})

	//标题
	w.OnHTML(" div > div.col-left > div.content-wrap > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//摘要
	w.OnHTML("div > div.col-left > div.content-wrap > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})

	//作者
	w.OnHTML(" div.container > div > div.col-left > div.content-wrap > span.author-info.view > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//正文
	w.OnHTML("section.article-detail.news-wrap.research-detail.small-font", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

}
