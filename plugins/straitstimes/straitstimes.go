package straitstimes

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("straitstimes", "海峡时报", "https://www.straitstimes.com/")

	w.SetStartingUrls([]string{"https://www.straitstimes.com/singapore/jobs",
		"https://www.straitstimes.com/singapore/housing",
		"https://www.straitstimes.com/singapore/parenting-education",
		"https://www.straitstimes.com/singapore/politics",
		"https://www.straitstimes.com/singapore/health",
		"https://www.straitstimes.com/singapore/transport",
		"https://www.straitstimes.com/singapore/courts-crime",
		"https://www.straitstimes.com/singapore/consumer",
		"https://www.straitstimes.com/singapore/environment",
		"https://www.straitstimes.com/singapore/community",
		"https://www.straitstimes.com/asia/se-asia",
		"https://www.straitstimes.com/asia/east-asia",
		"https://www.straitstimes.com/asia/south-asia",
		"https://www.straitstimes.com/asia/australianz",
		"https://www.straitstimes.com/world/united-states",
		"https://www.straitstimes.com/world/europe",
		"https://www.straitstimes.com/world/middle-east",
		"https://www.straitstimes.com/opinion/st-editorial",
		"https://www.straitstimes.com/opinion/forum",
		"https://www.straitstimes.com/business/economy",
		"https://www.straitstimes.com/business/invest",
		"https://www.straitstimes.com/business/banking",
		"https://www.straitstimes.com/business/companies-markets",
		"https://www.straitstimes.com/business/property",
		"https://www.straitstimes.com/tech/tech-news"})

	// 从index访问新闻
	w.OnHTML("a.stretched-link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	// report.title
	w.OnHTML("h1.headline", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// report .content
	w.OnHTML("div.ds-field-items > div > div:nth-child(1) > div", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	//report.publish time
	w.OnHTML("div.story-changeddate", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.author
	w.OnHTML("div.group-info", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
}
