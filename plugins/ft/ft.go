package ft

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("ft", "每日金融时报", "https://www.ft.lk/")

	w.SetStartingUrls([]string{"https://www.ft.lk/news/56",
		"https://www.ft.lk/education/10515",
		"https://www.ft.lk/entrepreneurship/41",
		"https://www.ft.lk/financial-services/42",
		"https://www.ft.lk/youthcareershigher-education/30",
		"https://www.ft.lk/hr/47",
		"https://www.ft.lk/other-sectors/57",
		"https://www.ft.lk/business/34",
		"https://www.ft.lk/opinion/14",
		"https://www.ft.lk/healthcare/45",
		"https://www.ft.lk/special-report/22",
		"https://www.ft.lk/sustainability__environment/10519",
		"https://www.ft.lk/it-telecom-tech/50",
		"https://www.ft.lk/shippingaviation/21",
		"https://www.ft.lk/harmony_page/10523",
		"https://www.ft.lk/special-editions",
		"https://www.ft.lk/advertorial/10529",
		"https://www.ft.lk/energy/10509",
		"https://www.ft.lk/in-depth/48",
		"https://www.ft.lk/markets/10518",
		"https://www.ft.lk/motor/55",
		"https://www.ft.lk/editorial/58",
		"https://www.ft.lk/ft-click/15",
		"https://www.ft.lk/agriculture/31",
		"https://www.ft.lk/international/49"})

	// 从翻页器获取链接并访问
	w.OnHTML(".pagination>li>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})
	// 从index访问新闻
	w.OnHTML("a.date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML(".col-md-6>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	// report.title
	w.OnHTML(".main>div>div>header", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// report .content
	w.OnHTML("header.inner-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	//report.publish time
	w.OnHTML(".gtime", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
}
