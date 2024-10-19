package storage

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1753", "罗彻斯特理工学院", "https://www.rit.edu/")

	engine.SetStartingURLs([]string{"https://www.rit.edu/sitemap.xml"})

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
		engine.Visit(element.Text, crawlers.Index)
		if strings.Contains(element.Text, "/news/") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".col-lg-8.pr-lg-5", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
