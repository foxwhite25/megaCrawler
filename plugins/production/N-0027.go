package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0027", "Calculated Risk Blog", "https://www.calculatedriskblog.com/")

	engine.SetStartingURLs([]string{"https://www.calculatedriskblog.com/sitemap.xml"})

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
		if strings.Contains(element.Text, "/sitemap.xml?") {
			engine.Visit(element.Text, crawlers.Index)
		} else {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("h2.date-header", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.ReplaceAll(element.Text, "Updated ", "")
	})

	engine.OnHTML("div.post-body.entry-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("script, span.fn").Remove()

		directText := element.DOM.Text()
		ctx.Content += strings.ReplaceAll(directText, "\n", "")
	})
}
