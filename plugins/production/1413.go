package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1413", "新闻秘书办公室", "https://www.pna.gov.ph/")

	engine.SetStartingURLs([]string{
		"https://www.pna.gov.ph/categories/national",
		"https://www.pna.gov.ph/categories/sona-2024",
		"https://www.pna.gov.ph/categories/bagong-pilipinas",
		"https://www.pna.gov.ph/categories/provincial",
		"https://www.pna.gov.ph/categories/business",
		"https://www.pna.gov.ph/categories/foreign",
		"https://www.pna.gov.ph/categories/sports",
		"https://www.pna.gov.ph/categories/travel-and-tourism",
		"https://www.pna.gov.ph/categories/health-and-lifestyle",
		"https://www.pna.gov.ph/categories/features",
		"https://www.pna.gov.ph/categories/events",
		"https://www.pna.gov.ph/categories/science-and-technology",
		"https://www.pna.gov.ph/categories/arts-and-entertainment",
		"https://www.pna.gov.ph/categories/media-security",
		"https://www.pna.gov.ph/categories/foi",
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

	engine.OnHTML(".flex-1 > h2 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".-space-x-px > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".article-content.mt-8.prose.max-w-none > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
