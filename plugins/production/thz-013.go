package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("thz-013", "CruiseGuide & CruiseGuideOnline.com", "https://www.cruisewatch.com/cruise-guide")

	engine.SetStartingURLs([]string{"https://www.cruisewatch.com/cruise-guide/cruise-costs",
		"https://www.cruisewatch.com/cruise-guide/plan-your-cruise",
		"https://www.cruisewatch.com/cruise-guide/short-cruises",
		"https://www.cruisewatch.com/cruise-guide/cruise-food",
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

	engine.OnHTML(".port-description > a:nth-child(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".thypography-content>.text-block + .text-block  p,span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find(".thypography-content>.text-block + .text-block  p > a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
