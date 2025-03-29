package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("3401", "Beyond Bulls and Bears", "https://global.beyondbullsandbears.com/")

	engine.SetStartingURLs([]string{
		"https://global.beyondbullsandbears.com/category/alternatives/",
		"https://global.beyondbullsandbears.com/category/equity/",
		"https://global.beyondbullsandbears.com/category/fixed-income/",
		"https://global.beyondbullsandbears.com/category/multi-asset/",
		"https://global.beyondbullsandbears.com/category/emerging-markets/",
		"https://global.beyondbullsandbears.com/category/perspectives/",
	})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("div.col-sm-6.col-md-6 > div.card > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("div.nav-previous > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("div.article-intro > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description += element.Text
	})

	engine.OnHTML("div.media-body.author-body > p > strong", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	engine.OnHTML("div.entry-content > p,div.entry-content > h2,div.entry-content > p,div.entry-content > ul",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})

}
