package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("GQT0015", "RT: For Decision Makers in Respiratory Care", "https://respiratory-therapy.com/")

	engine.SetStartingURLs([]string{
		"https://respiratory-therapy.com/disorders-diseases/chronic-pulmonary-disorders/",
		"https://respiratory-therapy.com/disorders-diseases/infectious-diseases/",
		"https://respiratory-therapy.com/resource-center/webinars/",
		"https://respiratory-therapy.com/disorders-diseases/critical-care/",
		"https://respiratory-therapy.com/disorders-diseases/cardiopulmonary-thoracic/",
		"https://respiratory-therapy.com/disorders-diseases/sleep-medicine/",
		"https://respiratory-therapy.com/public-health/smoking/",
		"https://respiratory-therapy.com/public-health/healthcare-policy/",
		"https://respiratory-therapy.com/public-health/pediatrics/",
		"https://respiratory-therapy.com/products-treatment/monitoring-treatment/"})

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

	engine.OnHTML(".entry-title>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".post-meta.vcard>p>span:first-child", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
	engine.OnHTML(".post-meta.vcard>p>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	engine.OnHTML(".post-content.entry-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	engine.OnHTML(".next.page-numbers", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
