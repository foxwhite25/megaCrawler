package storage

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1771", "Pressenza International Press Agency", "https://www.pressenza.com/")

	engine.SetStartingURLs([]string{"https://www.pressenza.com/wp-sitemap.xml"})

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
		if strings.Contains(element.Request.URL.String(), "wp-sitemap.xml") {
			if strings.Contains(element.Text, "wp-sitemap-posts-post") {
				ctx.Visit(element.Text, crawlers.Index)
				return
			}
			return
		}
		ctx.Visit(element.Text, crawlers.News)
	})
}
