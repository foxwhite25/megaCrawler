package dev

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1482", "澳大利亚国库部", "https://treasury.gov.au/")

	engine.SetStartingURLs([]string{
		"https://treasury.gov.au/sitemap.xml?page=1",
		"https://treasury.gov.au/sitemap.xml?page=2",
		"https://treasury.gov.au/sitemap.xml?page=3",
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
		if strings.Contains(element.Text, "/media-release/") || strings.Contains(element.Text, "/event/") ||
			strings.Contains(element.Text, "/publication/") || strings.Contains(element.Text, "/speech/") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".odd > span > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fileURL := element.Attr("href")
		if strings.Contains(fileURL, ".pdf") {
			url, err := element.Request.URL.Parse(element.Attr("href"))
			if err != nil {
				crawlers.Sugar.Error(err.Error())
				return
			}
			ctx.File = append(ctx.File, url.String())
			ctx.PageType = crawlers.Report
		}
	})

	engine.OnHTML("time.datetime", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(".node__content > div > h2, .node__content > div > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
