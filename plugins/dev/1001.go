package dev

import (
	"github.com/gocolly/colly/v2"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

func init() {
	engine := crawlers.Register("1001", "Carnegie_europe", "https://carnegieeurope.eu/")

	engine.SetStartingURLs([]string{"https://carnegieendowment.org/__sitemap__/research.xml"})

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
		engine.Visit(element.Text, crawlers.Report)
	})

	engine.OnHTML(".body", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = crawlers.StandardizeSpaces(element.Text)
	})
}

//package dev
// import (
// 	"strings"

// 	"megaCrawler/crawlers"
// 	"megaCrawler/extractors"

// 	"github.com/gocolly/colly/v2"
// )

// func init() {
// 	engine := crawlers.Register("1001", "Carnegie_europe", "https://carnegieeurope.eu/")

// 	engine.SetStartingURLs([]string{"https://carnegieendowment.org/research?lang=en"})

// 	extractorConfig := extractors.Config{
// 		Author:       true,
// 		Image:        true,
// 		Language:     true,
// 		PublishDate:  true,
// 		Tags:         true,
// 		Text:         true,
// 		Title:        true,
// 		TextLanguage: "",
// 	}

// 	extractorConfig.Apply(engine)

// 	engine.OnXML("//href", func(element *colly.XMLElement, ctx *crawlers.Context) {
// 		switch {
// 		case strings.Contains(ctx.URL, "expert"):
// 			engine.Visit(element.Text, crawlers.Expert)
// 		case strings.Contains(ctx.URL, "reasearch"):
// 			engine.Visit(element.Text, crawlers.News)
// 		}
// 	})
// }
