package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0055", "Nokiapoweruser", "https://nokiapoweruser.com/")

	engine.SetStartingURLs([]string{"https://nokiapoweruser.com/sitemap-index-1.xml"})

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
		if strings.Contains(element.Text, "sitemap") {
			engine.Visit(element.Text, crawlers.Index)
		} else if !strings.Contains(element.Text, "sitemap") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".td-post-content > p[style=\"text-align:justify;\"], .td-post-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
