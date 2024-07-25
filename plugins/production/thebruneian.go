package production

import (
	"time"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("thebruneian", "thebruneian", "https://thebruneian.news/")

	w.SetStartingURLs([]string{
		"https://thebruneian.news/",
	})

	w.OnHTML(".author-name", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("div.entry-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML("ul.mobile-menu>li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("a.page-numbers", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("a.evo-overlay-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML(".evo-entry-title>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("div.evo-post-date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		now, done := crawlers.ParseRelativeTime(element.Text)
		if done {
			return
		}
		ctx.PublicationTime = now.Format(time.RFC3339)
	})
	w.OnHTML("h1.evo-entry-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})
}
