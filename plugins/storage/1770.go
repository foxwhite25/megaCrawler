package storage

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1770", "国家先驱论坛报", "https://www.nationalheraldindia.com/")

	engine.SetStartingURLs([]string{"https://www.nationalheraldindia.com/sitemap.xml"})

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
		if strings.Contains(element.Request.URL.String(), "sitemap.xml") {
			if strings.Contains(element.Text, "daily") {
				ctx.Visit(element.Text, crawlers.Index)
				return
			}
			return
		}
		ctx.Visit(element.Text, crawlers.News)
	})

	// engine.OnHTML(".styles-m__story-content__1xPvO", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	ctx.Content += element.Text
	// })
}
