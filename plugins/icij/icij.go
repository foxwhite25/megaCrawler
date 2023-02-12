package icij

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("icij", "国际调查记者联盟", "https://www.icij.org/")

	w.SetStartingUrls([]string{"https://www.icij.org/investigations/trafficking-inc/",
		"https://www.icij.org/investigations/uber-files/",
		"https://www.icij.org/investigations/pandora-papers/",
		"https://www.icij.org/investigations/panama-papers/",
		"https://www.icij.org/investigations/",
		"https://www.icij.org/inside-icij/",
		"https://www.icij.org/tags/data-journalism/",
		"https://www.icij.org/tags/impact/",
		"https://www.icij.org/tags/accountability/",
		"https://www.icij.org/tags/offshore-secrecy/",
		"https://www.icij.org/tags/europe/",
		"https://www.icij.org/tags/tax-havens/",
		"https://www.icij.org/tags/data-journalism/",
		"https://www.icij.org/tags/investigative-reporting/",
		"https://www.icij.org/tags/icij-members/",
		"https://www.icij.org/tags/investigative-journalism/",
		"https://www.icij.org/tags/africa/",
		"https://www.icij.org/tags/press-freedom/"})

	// 从翻页器获取链接并访问
	w.OnHTML("#icijorg > div.wrap.container.mb-5 > div > div.col-md-9.col-12 > div > div > div > article > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从index访问新闻
	w.OnHTML("#icijorg > div:nth-child(6) > div > div > div:nth-child(2) > div > article > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML("#icijorg > div.wrap.container.mb-3 > div > div > div > div > article > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML("#icijorg > div:nth-child(5) > div.wrap.container.mb-3 > div > div > div:nth-child(2) > div > article > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML("div.col-12.secondary-featured-articles > div > div:nth-child(1) > article > div> div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML("#icijorg > div:nth-child(6) > div > div.col-12.col-md-9.mb-3 > div > div > div:nth-child(2) > div> article > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML(".article-title__title>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	//
	// report.title
	w.OnHTML("h1.entry-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})
	// report.description
	w.OnHTML("div.post-excerpt>p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})
	// report.author
	w.OnHTML("a.post-author", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//report.publish time
	w.OnHTML("time.post-published", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report .content
	w.OnHTML("div.post-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	//内含Expert
	w.OnHTML("a.post-author", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})
	// expert.Name
	w.OnHTML("#user-content>h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})
	// expert.description
	w.OnHTML("div.post-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
}
