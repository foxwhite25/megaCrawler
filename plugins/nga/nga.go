package nga

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("nga", "国家地理空间情报局", "https://www.nga.mil/")
	w.SetStartingUrls([]string{"https://www.nga.mil/news/News.html"})

	//index
	w.OnHTML("span.usa-pagination__link-text", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问新闻
	w.OnHTML("a.usa-button", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//获取标题
	w.OnHTML("h1.font-family-heading.text-white", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//获取tag(忽略可能出现的邮件地址)
	w.OnHTML(" div > div> div > p > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".mil") {

		} else {
			ctx.Tags = append(ctx.Tags, element.Text)
		}
	})

	//获取正文
	w.OnHTML(".margin-bottom-4>p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content += element.Text
	})
}
