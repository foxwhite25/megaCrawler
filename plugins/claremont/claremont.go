package claremont

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("claremont", "克萊蒙研究所", "https://www.claremont.org/")

	w.SetStartingUrls([]string{"https://www.claremont.org/scholars/", "https://www.claremont.org/media/"})

	//内含Expert
	w.OnHTML("div.scholars_inner>div>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})
	// expert.Name
	w.OnHTML("div.detail_bio>h2.content_header:nth-child(1)", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})
	// expert.description
	w.OnHTML("div.detail_bio>article.listing_article", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
}
