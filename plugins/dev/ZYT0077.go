package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("ZYT0077", "bpi", "https://www.bpi.com.ph")

	engine.SetStartingURLs([]string{"https://www.bpi.com.ph/sitemap.xml"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  false,
		Tags:         false,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.SetTimeout(60 * time.Second)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "news") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".text>div>div>p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".info__editorial__timings>span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	engine.OnHTML(".bread-crum-container.custom-container>div:nth-child(1)>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.FirstChannel = element.Text
	})

	engine.OnHTML(".bread-crum-container.custom-container>div:nth-child(2)>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SecondChannel = element.Text
	})

	engine.OnHTML(".d-flex.flex-column>span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})
}
