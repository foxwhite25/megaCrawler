package production

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1730", "伊朗新闻电视台", "https://www.presstv.ir/")

	engine.SetStartingURLs([]string{"https://www.presstv.ir/sitemap.xml"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "sitemap") {
			engine.Visit(element.Text, crawlers.Index)
			return
		}
		engine.Visit(element.Text, crawlers.News)
	})
}
