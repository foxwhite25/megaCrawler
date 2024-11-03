package storage

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1774", "Tehelka", "https://tehelkahindi.com/")

	engine.SetStartingURLs([]string{"https://tehelkahindi.com/wp-sitemap.xml"})

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
		switch {
		case strings.Contains(element.Text, "wp-sitemap-posts-post"):
			engine.Visit(element.Text, crawlers.Index)
		default:
			engine.Visit(element.Text, crawlers.News)
		}
	})
	engine.OnHTML(".td-post-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
