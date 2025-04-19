package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0023", "BrandlandUSA", "https://www.brandlandusa.com/")

	engine.SetStartingURLs([]string{
		"https://www.brandlandusa.com/post-sitemap.xml",
		"https://www.brandlandusa.com/post-sitemap2.xml",
	})

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
		engine.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML(".entry-the-content > p, .entry-the-content > blockquote, .entry-the-content > ul",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			element.DOM.Find("script").Remove()
			element.DOM.Find("noscript").Remove()

			directText := element.DOM.Text()
			ctx.Content += strings.Join(strings.Fields(directText), " ")
		})
}
