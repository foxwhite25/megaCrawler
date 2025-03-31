package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2416", "Furniture Today", "https://www.furnituretoday.com/")

	engine.SetStartingURLs([]string{
		"https://www.furnituretoday.com/category/las-vegas-market-furniture-show/",
		"https://www.furnituretoday.com/category/technology",
		"https://www.furnituretoday.com/category/supply-chain/",
		"https://www.furnituretoday.com/category/e-commerce/",
		"https://www.furnituretoday.com/category/furniture-retailing/",
		"https://www.furnituretoday.com/category/furniture-people/",
		"https://www.furnituretoday.com/category/manufacturers/",
		"https://www.furnituretoday.com/category/business-news/",
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

	engine.OnHTML(".content-section > h2 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("div.wp-pagenavi > a.nextpostslink", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("div.news-detail > p,div.news-detail > ul", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
