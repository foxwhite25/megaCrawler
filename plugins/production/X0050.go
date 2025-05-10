package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0050", "新喀里多尼亚第一频道", "https://www.francetvinfo.fr/")

	engine.SetStartingURLs([]string{"https://www.francetvinfo.fr/sitemap_index.xml"})

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
		if strings.Contains(element.Text, "article") {
			engine.Visit(element.Text, crawlers.Index)
		} else if (!strings.Contains(element.Text, "xml")) && (!strings.Contains(element.Text, "sports")) {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".c-body > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
