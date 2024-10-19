package production

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1018", "新美国安全中心", "https://www.cnas.org/")

	engine.SetStartingURLs([]string{"https://www.cnas.org/sitemaps-1-sitemap.xml"})

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
		if strings.Contains(element.Text, "section-articles") {
			engine.Visit(element.Text, crawlers.Index)
			return
		}
		if strings.Contains(element.Request.URL.String(), "section-articles") {
			engine.Visit(element.Text, crawlers.News)
			return
		}
	})
}
