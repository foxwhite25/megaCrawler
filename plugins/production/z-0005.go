package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("z-0005", "Supreme Court of the Philippines", "https://sc.judiciary.gov.ph/")

	engine.SetStartingURLs([]string{"https://sc.judiciary.gov.ph/press-releases/"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)
	engine.OnHTML(".thumb > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML("#content > div > div > section.elementor-section.elementor-top-section.elementor-element.elementor-element-0a51d59.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div  p,.elementor-widget-container > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
