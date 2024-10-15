package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1021", "阿彭斯研究所", "https://www.aspeninstitute.org/")

	engine.SetStartingURLs([]string{"https://www.aspeninstitute.org/sitemap_index.xml"})

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
		if strings.Contains(element.Text, "sitemap") {
			engine.Visit(element.Text, crawlers.Index)
			return
		}

		time.Sleep(400 * time.Millisecond) //每次访问延迟400ms

		engine.Visit(element.Text, crawlers.News)
	})
}
