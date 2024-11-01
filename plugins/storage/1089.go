package storage

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1089", "每日星报", "https://www.dailystar.co.uk")

	engine.SetStartingURLs([]string{"https://www.dailystar.co.uk/sitemaps/sitemap_index.xml"})

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
		switch {

		case strings.Contains(element.Text, "sitemap"):
			engine.Visit(element.Text, crawlers.Index)

		default:
			engine.Visit(element.Text, crawlers.News)
		}
	})

}
