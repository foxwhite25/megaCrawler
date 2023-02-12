package iol

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("iol", "比勒陀利亚新闻", "https://www.iol.co.za/")

	w.SetStartingUrls([]string{"https://www.iol.co.za/news/south-africa/northern-cape",
		"https://www.iol.co.za/personal-finance/financial-planning",
		"https://www.iol.co.za/business-report/brics",
		"https://www.iol.co.za/property/commercial",
		"https://www.iol.co.za/opinion",
		"https://www.iol.co.za/news/africa",
		"https://www.iol.co.za/business-report/opinion",
		"https://www.iol.co.za/news/south-africa/limpopo",
		"https://www.iol.co.za/news/crime-and-courts",
		"https://www.iol.co.za/news/world",
		"https://www.iol.co.za/business-report/careers",
		"https://www.iol.co.za/news/south-africa/free-state",
		"https://www.iol.co.za/personal-finance/tax",
		"https://www.iol.co.za/technology/internet-of-things",
		"https://www.iol.co.za/business-report/budget",
		"https://www.iol.co.za/technology/fintech",
		"https://www.iol.co.za/lifestyle/health",
		"https://www.iol.co.za/news/south-africa/north-west",
		"https://www.iol.co.za/news/politics",
		"https://www.iol.co.za/news/south-africa/mpumalanga",
		"https://www.iol.co.za/personal-finance/investments",
		"https://www.iol.co.za/business-report/ending-poverty-in-china",
		"https://www.iol.co.za/business-report/economy",
		"https://www.iol.co.za/news/environment",
		"https://www.iol.co.za/business-report/markets",
		"https://www.iol.co.za/personal-finance/debt",
		"https://www.iol.co.za/technology/gadgets",
		"https://www.iol.co.za/news/south-africa/western-cape",
		"https://www.iol.co.za/personal-finance/retirement",
		"https://www.iol.co.za/news/traffic",
		"https://www.iol.co.za/technology/mobile",
		"https://www.iol.co.za/business-report/entrepreneurs",
		"https://www.iol.co.za/property/residential",
		"https://www.iol.co.za/video",
		"https://www.iol.co.za/news/south-africa/kwazulu-natal",
		"https://www.iol.co.za/news/south-africa/eastern-cape",
		"https://www.iol.co.za/business-report/international",
		"https://www.iol.co.za/news/south-africa/gauteng"})

	// 从index访问新闻
	w.OnHTML(".Link__StyledLink-sc-1pz26op-0", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	// report.title
	w.OnHTML("#root > main > article > div.sc-kkGfuU.iYTFhX > div > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// report .content
	w.OnHTML("#root > main > article > div.sc-kkGfuU.cxerjY > div.sc-kkGfuU.taSns", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	//report.publish time
	w.OnHTML("#root > main > article > div.sc-kkGfuU.iYTFhX > div > div.sc-kkGfuU.gwzeLU > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.author
	w.OnHTML("#root > main > article > div.sc-kkGfuU.iYTFhX > aside.sc-EHOje.escQeg > div > div > p.sc-cIShpX.qbOpm > a > span", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

}
