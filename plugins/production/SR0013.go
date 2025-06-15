package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("SR0013", "SHAWN E TUMA - Blog", "https://shawnetuma.com/")

	engine.SetStartingURLs([]string{"https://shawnetuma.com/sitemap-1.xml"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
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

	engine.OnHTML("div.entry-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})
}
