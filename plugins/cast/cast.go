package cast

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("cast", "战略与技术分析中心", "http://cast.ru/eng/")
	w.SetStartingUrls([]string{"http://cast.ru/eng/"})

	//index
	w.OnHTML(" div.middle > div > main > div> ul > li > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问news
	w.OnHTML("body > div.wrapper > div.middle > div > main > div.double-block.clearfix > div > ul > li> div.body > div> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML("div.item-body > div.item-name > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//获取标题
	w.OnHTML("div.middle > div > main > div.article.clearfix > div > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//获取正文
	w.OnHTML("body > div.wrapper > div.middle > div > main > div.article.clearfix > div", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
}
