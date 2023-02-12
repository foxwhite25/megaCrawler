package foreignaffairs

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("foreignaffairs", "外事网/外交/外交事务", "https://www.foreignaffairs.com/")

	w.SetStartingUrls([]string{"https://www.foreignaffairs.com/browse/view-all",
		"https://www.foreignaffairs.com/browse/review",
		"https://www.foreignaffairs.com/podcasts/foreign-affairs-interview",
		"https://www.foreignaffairs.com/anthologies",
		"https://www.foreignaffairs.com/events",
		"https://www.foreignaffairs.com/topics/biden-administration",
		"https://www.foreignaffairs.com/tags/war-ukraine",
		"https://www.foreignaffairs.com/tags/coronavirus",
		"https://www.foreignaffairs.com/topics/climate-change",
		"https://www.foreignaffairs.com/topics/cybersecurity",
		"https://www.foreignaffairs.com/tags/nationalism",
		"https://www.foreignaffairs.com/topics/democratization",
		"https://www.foreignaffairs.com/topics/economics",
		"https://www.foreignaffairs.com/topics/globalization",
		"https://www.foreignaffairs.com/topics/refugees-migration",
		"https://www.foreignaffairs.com/topics/us-foreign-policy",
		"https://www.foreignaffairs.com/topics/war-military-strategy",
		"https://www.foreignaffairs.com/regions/united-states",
		"https://www.foreignaffairs.com/regions/ukraine",
		"https://www.foreignaffairs.com/regions/russian-federation",
		"https://www.foreignaffairs.com/regions/china",
		"https://www.foreignaffairs.com/regions/iran",
		"https://www.foreignaffairs.com/regions/north-korea",
		"https://www.foreignaffairs.com/regions/afghanistan",
		"https://www.foreignaffairs.com/regions/ethiopia",
		"https://www.foreignaffairs.com/regions",
		"https://www.foreignaffairs.com/browse/essay",
		"https://www.foreignaffairs.com/browse/snapshot",
		"https://www.foreignaffairs.com/browse/articles-with-audio",
		"https://www.foreignaffairs.com/browse/review-essay",
		"https://www.foreignaffairs.com/browse/interview",
		"https://www.foreignaffairs.com/browse/response"})

	// 从index访问新闻
	w.OnHTML("a.ob-l", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML("a.event-teaser-header_link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML("a.magazine-list-item--image-link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	// report.title
	w.OnHTML("h1.article-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = ctx.Title + element.Text
	})
	w.OnHTML("h2.article-subtitle", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = ctx.Title + element.Text
	})
	w.OnHTML("h1.event-detail-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})
	// report.author
	w.OnHTML("a.article-byline-author", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	//report.publish time
	w.OnHTML("span.article-header--metadata-date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	w.OnHTML("time.event-detail-info-time", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report .content
	w.OnHTML("div.paywall-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	w.OnHTML("div.event-detail-container", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	//内含Expert
	w.OnHTML("a.authors-bio__link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})
	// expert.Name
	w.OnHTML("h1.title-author-name", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})
	// expert.description
	w.OnHTML("div.author-bio", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

}
