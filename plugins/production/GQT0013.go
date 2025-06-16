package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("GQT0013", "RollCall call", "https://rollcall.com/")

	engine.SetStartingURLs([]string{
		"https://rollcall.com/section/campaigns/",
		"https://rollcall.com/category/energy/",
		"https://rollcall.com/category/defense/",
		"https://rollcall.com/section/white-house/",
		"https://rollcall.com/section/congress/",
		"https://rollcall.com/category/health-care/",
		"https://rollcall.com/category/energy/",
		"https://rollcall.com/section/heard-on-the-hill/",
		"https://rollcall.com/category/fintech/",
		"https://rollcall.com/podcast/cq-budget-podcast/"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".title>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".byline-date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
	engine.OnHTML(".font-semibold+a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	engine.OnHTML(".font-lyon>p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	engine.OnHTML(".next.page-numbers", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
