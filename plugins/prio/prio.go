package prio

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("prio", "奥斯陆和平研究所", "https://www.prio.org/")
	w.SetStartingUrls([]string{"https://www.prio.org/people"})

	//访问专家
	w.OnHTML(" div.columns.is-multiline > div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//专家姓名
	w.OnHTML(".personal-info > div > h2", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	//专家头衔
	w.OnHTML(".personal-info > div > h4", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//专家描述
	w.OnHTML("card-content-padding content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})

	//专家领域
	w.OnHTML(" div.person-topics > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area += element.Text + " "
	})

	//访问report
	w.OnHTML(".publication  > div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	//访问news
	w.OnHTML(".news  > div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//文章标题
	w.OnHTML("div > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//时间
	w.OnHTML("p.subtitle", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//新闻正文
	w.OnHTML("article-body content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	//报告正文
	w.OnHTML(".content.mt-5 > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content += element.Text
	})

	//pdf
	w.OnHTML("a.download-button", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})

	//作者
	w.OnHTML("div.authors > div > div> a > div > div.card-content-padding > h5", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
}
