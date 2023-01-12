package claremontreviewofbooks

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("claremontreviewofbooks", "克萊蒙研究所", "https://claremontreviewofbooks.com/")

	w.SetStartingUrls([]string{"https://claremontreviewofbooks.com/",
		"https://claremontreviewofbooks.com/articles/essays/",
		"https://claremontreviewofbooks.com/articles/book-reviews/",
		"https://claremontreviewofbooks.com/podcast/",
		"https://claremontreviewofbooks.com/authors/"})

	// 从翻页器获取链接并访问
	w.OnHTML("div.comp__pagination-next>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从index访问新闻
	w.OnHTML("h1.post-title>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	// report.title
	w.OnHTML("h1.content-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})
	// report.description
	w.OnHTML("p.content-preview", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})
	// report.author
	w.OnHTML("div.box__name", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	// report .content
	w.OnHTML("div.content-restricted", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	w.OnHTML("div.col-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	//内含Expert
	w.OnHTML("article.post-type__author>h1.post-title>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})
	// expert.Name
	w.OnHTML("h1.single-author-meta-name", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})
	// expert.description
	w.OnHTML("div.single-author-bio>p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
}
