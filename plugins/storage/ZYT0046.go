package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("ZYT0046", "steelorbis", "https://www.steelorbis.com")

	engine.SetStartingURLs([]string{"https://www.steelorbis.com/steel-news/latest-news/"})

	extractorConfig := extractors.Config{
		Author:       false, // 文章无作者
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("a[style=\"text-decoration:none;\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".col-md-12.col-lg-8", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fulltextwithlocation := element.Text
		textwithlocation := strings.Split(fulltextwithlocation, "|")
		locationpart := strings.TrimSpace(textwithlocation[1])
		locationpart = strings.ReplaceAll(locationpart, "fa-solid fa-map-pin", "")
		locationpart = strings.TrimSpace(locationpart)
		ctx.Location = locationpart
	})

	engine.OnHTML(".col-md-12.col-lg-8", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fulltextwithdate := element.Text
		textwithdate := strings.Split(fulltextwithdate, "|")
		datepart := strings.TrimSpace(textwithdate[0])
		datepart = strings.ReplaceAll(datepart, "fa-regular fa-clock", "")
		datepart = strings.TrimSpace(datepart)
		ctx.PublicationTime = datepart
	})

	engine.OnHTML("#contentDiv", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".active + li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
