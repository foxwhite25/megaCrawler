package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("aa", "aa", "https://www.aa.com.tr/en")

	w.SetStartingURLs([]string{
		"https://www.aa.com.tr/en",
	})

	w.OnHTML("body > div:nth-child(3) > main > div > div.detay-bg > div > div > div > span:nth-child(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML(" span.fa.fa-user > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML(".detay-icerik", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML(".article-post-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML("a.text-white", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML(".borderR>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML(".col-sm-8>h3>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML(".col-sm-12>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("span.tarih", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(strings.Split(element.Text, " - ")[0])
	})
	w.OnHTML(".detay-spot-category>h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})
}
