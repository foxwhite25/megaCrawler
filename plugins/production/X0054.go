package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0054", "nextbigfuture", "https://www.nextbigfuture.com/")

	engine.SetStartingURLs([]string{"https://www.nextbigfuture.com/sitemap_index.xml"})

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
		} else if strings.Contains(element.Text, "/20") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".entry-content > p, .entry-content > i", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		element.DOM.Find("img").Remove()
		element.DOM.Find("iframe").Remove()
		element.DOM.Find("script").Remove()
		element.DOM.Find("noscript").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
