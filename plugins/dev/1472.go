package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1472", "澳大利亚教育与培训部", "https://www.education.gov.au/")

	engine.SetStartingURLs([]string{
		"https://www.education.gov.au/sitemap.xml?page=1",
	})

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
		if strings.Contains(element.Text, "/announcements/") || strings.Contains(element.Text, "/newsroom/articles/") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	//公告和新闻不同的selector
	engine.OnHTML(".node__content.clearfix > div > time, .container > time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(".node__content.clearfix > header > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description += element.Text
	})

	engine.OnHTML(`.paragraph.paragraph--type--text.paragraph--view-mode--default > div > p,
	.paragraph.paragraph--type--text.paragraph--view-mode--default > div > ul`,
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
