package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0012", "皮马社区学院", "https://www.pima.edu/")

	engine.SetStartingURLs([]string{"https://www.pima.edu/sitemap.xml"})

	// engine.SetParallelism(2)

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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "/news/stories/") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".col-12.col-md-9.wysiwyg div:not(.link-callout)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
