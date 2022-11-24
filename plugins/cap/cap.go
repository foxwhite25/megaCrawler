package cap

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("cna", "美国进步中心", "https://www.americanprogress.org/")
	w.SetStartingUrls([]string{"https://www.americanprogress.org/issues/#topics", "https://www.americanprogress.org/experts/"})

	//访问人物
	w.OnHTML(" div.people1 > div.-o\\:h > div > article> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//人物名称
	w.OnHTML(" div > header > div.header1-main > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Expert {
			ctx.Name = element.Text
		}
	})

	//领域
	w.OnHTML(" div > header > div.header1-main > h3", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area += element.Text + " "
	})

	//专家介绍
	w.OnHTML(" div > div.-xw\\:4.-mx\\:a.-mb\\:3 > div > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description += element.Text
	})

	//专家图片
	w.OnHTML("div.header1-side > figure > span.img1.-top > img", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Image = []string{element.Attr("href")}
	})

	//index
	w.OnHTML(" div:nth-child(1) > div > div.-o\\:h > div > div", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})
	w.OnHTML("#showmore > div > ul > li> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})
	w.OnHTML("div.archives1-wrap > p > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//从index访问report
	w.OnHTML(" div.archives1-wrap > div > article > div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	//标题
	w.OnHTML(" header > div.header2-wrap > div.header2-main > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//报告摘要
	w.OnHTML(" div > header > div.header2-wrap > div.header2-main > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description += element.Text
	})

	//作者
	w.OnHTML(".authors1-list span > span", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	//时间
	w.OnHTML(" div > header > div.header2-brow > time", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	//类型
	w.OnHTML(" div > header > div.header2-brow > span > span", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = element.Text
	})

	//正文
	w.OnHTML(".article1-main", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

}
