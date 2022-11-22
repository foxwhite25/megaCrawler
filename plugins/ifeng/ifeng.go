package ifeng

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("ifeng", "凤凰网国际智库", "https://news.ifeng.com/")
	w.SetStartingUrls([]string{"https://news.ifeng.com/shanklist/3-245389-/", "https://news.ifeng.com/c/special/81iPSl49iNc"})

	//访问新闻
	w.OnHTML(".modules-2CRmiqkc > div > div> div > div > div> p> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	w.OnHTML(" div.content-2VFX2Hqk > div > ul > li> div > h2 > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//获取标题
	w.OnHTML("h1.topic-z1ycr_72", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//获取时间
	w.OnHTML(".timeBref-2lHnksft", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//获取正文
	w.OnHTML("div.text-3w2e3DBc > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content += element.Text
	})
}
