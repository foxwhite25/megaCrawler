package brecorder

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("brecorder", "财经日报", "https://www.brecorder.com/")

	w.SetStartingUrls([]string{"https://www.brecorder.com/latest-news",
		"https://www.brecorder.com/br-research/analysis-and-comments",
		"https://www.brecorder.com/br-research/banking",
		"https://www.brecorder.com/br-research/bonds",
		"https://www.brecorder.com/br-research/chemicals",
		"https://www.brecorder.com/br-research/cotton-analysis",
		"https://www.brecorder.com/br-research/engineering",
		"https://www.brecorder.com/br-research/industry",
		"https://www.brecorder.com/br-research/interviews",
		"https://www.brecorder.com/br-research/miscellaneous",
		"https://www.brecorder.com/br-research/oil-and-gas",
		"https://www.brecorder.com/br-research/pharmaceuticals",
		"https://www.brecorder.com/br-research/power-generation",
		"https://www.brecorder.com/br-research/stocks",
		"https://www.brecorder.com/br-research/textile",
		"https://www.brecorder.com/br-research/transport",
		"https://www.brecorder.com/markets/forex",
		"https://www.brecorder.com/markets/mutual-funds",
		"https://www.brecorder.com/markets/financial",
		"https://www.brecorder.com/markets/paper",
		"https://www.brecorder.com/business-finance/companies",
		"https://www.brecorder.com/business-finance/industry",
		"https://www.brecorder.com/business-finance/taxes",
		"https://www.brecorder.com/business-finance/budgets",
		"https://www.brecorder.com/business-finance/interest-rates",
		"https://www.brecorder.com/pakistan/industries-sectors",
		"https://www.brecorder.com/business-finance/documents",
		"https://www.brecorder.com/pakistan/business-economy",
		"https://www.brecorder.com/world/china",
		"https://www.brecorder.com/business-finance/pakistan",
		"https://www.brecorder.com/business-finance/money-banking",
		"https://www.brecorder.com/world/mena",
		"https://www.brecorder.com/markets/stocks",
		"https://www.brecorder.com/markets/grains",
		"https://www.brecorder.com/world/middle-east",
		"https://www.brecorder.com/business-finance/managed-funds",
		"https://www.brecorder.com/world/south-asia",
		"https://www.brecorder.com/technology/it-and-computers",
		"https://www.brecorder.com/world/asia",
		"https://www.brecorder.com/markets/energy",
		"https://www.brecorder.com/pakistan/markets",
		"https://www.brecorder.com/world/usa",
		"https://www.brecorder.com/business-finance/real-estate",
		"https://www.brecorder.com/world/africa",
		"https://www.brecorder.com/business-finance/agriculture-and-allied",
		"https://www.brecorder.com/world/europe",
		"https://www.brecorder.com/perspectives",
		"https://www.brecorder.com/editorials",
		"https://www.brecorder.com/pakistan/elections",
		"https://www.brecorder.com/markets/yarn-prices",
		"https://www.brecorder.com/opinion",
		"https://www.brecorder.com/business-finance/general-news",
		"https://www.brecorder.com/markets/cotton-textile",
		"https://www.brecorder.com/pakistan/politics-policy"})

	// 从index访问新闻
	w.OnHTML("a.story__link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	// report.title
	w.OnHTML("h1.story__title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// report .content
	w.OnHTML(".story__content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = ctx.Content + element.Text
	})
	// report.author
	w.OnHTML(".story__byline", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	//report.publish time
	w.OnHTML(".timestamp--date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

}
