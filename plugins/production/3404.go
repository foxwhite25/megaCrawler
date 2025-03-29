﻿package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("3404", "BioSpace", "https://www.biospace.com/")

	engine.SetStartingURLs([]string{"https://www.biospace.com/sitemap.xml"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, ".xml") {
			engine.Visit(element.Text, crawlers.Index)
		} else {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("div.RichTextArticleBody.RichTextBody p, div.RichTextArticleBody.RichTextBody > ul",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			element.DOM.Find("script").Remove()
			directText := element.DOM.Text()
			ctx.Content += directText
		})
}
