package peoplesreview

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("peoplesreview", "人民评论", "https://www.peoplesreview.com.np/")

	w.SetStartingUrls([]string{"https://www.peoplesreview.com.np/category/opinion/from-far-near/",
		"https://www.peoplesreview.com.np/category/opinion/on-off-the-record/",
		"https://www.peoplesreview.com.np/category/opinion/nepali-netbook/",
		"https://www.peoplesreview.com.np/category/business-and-corporate/",
		"https://www.peoplesreview.com.np/category/news/editorial/",
		"https://www.peoplesreview.com.np/category/opinion/readers-forum/",
		"https://www.peoplesreview.com.np/category/opinion/viewpoint/",
		"https://www.peoplesreview.com.np/category/opinion/babbles/",
		"https://www.peoplesreview.com.np/category/news/current-news/news-analysis/",
		"https://www.peoplesreview.com.np/category/interviews-commentary/"})
	// 从翻页器获取链接并访问
	w.OnHTML("a.page-numbers", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从index访问新闻
	w.OnHTML(".uk-link>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	// report.title
	w.OnHTML(".uk-heading-primary", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// report .content
	w.OnHTML(".uk-text-large", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = ctx.Content + element.Text
	})
	//report.publish time
	w.OnHTML("body > div.uk-container > div.uk-grid > div> div > li > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	// report.author
	w.OnHTML(" div.uk-text-large.uk-text-justify > p:nth-child(3) > strong", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
}
