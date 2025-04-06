package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0002", "菲华电视台", "https://www.chinoy.tv/")

	engine.SetStartingURLs([]string{"https://www.chinoy.tv/sitemap_index.xml"})

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
		if strings.Contains(element.Text, "post-sitemap") {
			engine.Visit(element.Text, crawlers.Index)
		} else if (!strings.Contains(element.Text, ".xml")) && (strings.Contains(element.Text, "-")) {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".wd-entry-content > p,.wd-entry-content > h3", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
