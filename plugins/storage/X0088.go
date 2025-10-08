package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0088", "GoTyme Bank", "https://www.gotyme.com.ph/")

	engine.SetStartingURLs([]string{"https://www.gotyme.com.ph/sitemap.xml"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        false,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "news/") {
			engine.Visit(element.Text, crawlers.News)
		} else if strings.Contains(element.Text, "stories/") {
			engine.Visit(element.Text, crawlers.Report)
		}
	})

	engine.OnHTML(".px-lg-4.px-3.mb-4 > h4", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	engine.OnHTML(".mx-2 + .text-body-tertiary", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
		ctx.Location = "Philippines"
	})

	engine.OnHTML("small[style=\"color: #4809E9; font-weight: bold;\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
		ctx.Authors = append(ctx.Authors, "GoTyme Bank")
	})

	engine.OnHTML("div[class=\"row\"] > div:nth-last-child(1)  div > :not(:not(p, ol))", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
