package dailynews

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("dailynews", "每日新闻", "https://dailynews.co.tz/")

	w.SetStartingUrls([]string{"https://dailynews.co.tz/category/tanzania/dodoma/",
		"https://dailynews.co.tz/category/tanzania/rural-tanzania/",
		"https://dailynews.co.tz/category/tanzania/zonal/",
		"https://dailynews.co.tz/category/tanzania/zanzibar/",
		"https://dailynews.co.tz/category/tanzania/religion/",
		"https://dailynews.co.tz/category/tanzania/tourism/",
		"https://dailynews.co.tz/category/tanzania/history/",
		"https://dailynews.co.tz/category/tanzania/opportunities/",
		"https://dailynews.co.tz/category/world/africa/",
		"https://dailynews.co.tz/category/world/america/",
		"https://dailynews.co.tz/category/world/asia/",
		"https://dailynews.co.tz/category/world/europe/",
		"https://dailynews.co.tz/category/politics/parliament/",
		"https://dailynews.co.tz/category/politics/diplomacy/",
		"https://dailynews.co.tz/category/politics/election/",
		"https://dailynews.co.tz/category/society-culture/culture/",
		"https://dailynews.co.tz/category/society-culture/arts/",
		"https://dailynews.co.tz/category/tanzania/history/",
		"https://dailynews.co.tz/category/society-culture/heritage/",
		"https://dailynews.co.tz/category/society-culture/travel/",
		"https://dailynews.co.tz/category/society-culture/food-drinks/",
		"https://dailynews.co.tz/category/society-culture/women/",
		"https://dailynews.co.tz/category/society-culture/literature/",
		"https://dailynews.co.tz/category/business/investment/",
		"https://dailynews.co.tz/category/business/economy/",
		"https://dailynews.co.tz/category/business/finance/",
		"https://dailynews.co.tz/category/science-tech/",
		"https://dailynews.co.tz/category/health/",
		"https://dailynews.co.tz/category/in-depth/",
		"https://dailynews.co.tz/category/opinions/commentaries/",
		"https://dailynews.co.tz/category/opinions/analysis/",
		"https://dailynews.co.tz/category/opinions/editorials/",
		"https://dailynews.co.tz/category/opinions/opd/",
		"https://dailynews.co.tz/category/extraction/gas/",
		"https://dailynews.co.tz/category/extraction/mining/",
		"https://dailynews.co.tz/category/extraction/oil/",
		"https://dailynews.co.tz/category/multimedia/podcast/"})

	// 从翻页器获取链接并访问
	w.OnHTML("span.last-page>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从index访问新闻
	w.OnHTML("h2.entry-title>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML("a.all-over-thumb-link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	// report.title
	w.OnHTML("h1.post-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// report .content
	w.OnHTML("div.entry-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	//report.publish time
	w.OnHTML("div.entry-header>div>span.date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.author
	w.OnHTML("a.author-name", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

}
