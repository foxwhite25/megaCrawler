package dev

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1783", "Agency Tunis Afrique Press", "https://www.tap.info.tn/en/")

	engine.SetStartingURLs([]string{
		"https://www.tap.info.tn/en/Portal-Politics",
		"https://www.tap.info.tn/en/portal%20-%20economy",
		"https://www.tap.info.tn/en/portal_sciences_technology_eng",
		"https://www.tap.info.tn/en/portal%20-%20society",
		"https://www.tap.info.tn/en/portal%20-%20regions",
		"https://www.tap.info.tn/en/portal%20-%20world",
	})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        false,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".NewsItemText > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(`.NewsItemHeadline`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 网站需登录才能查看完整文章
	engine.OnHTML(".NewsItemCaption", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML("[style=\"color:red;font-size:22px;\"] + a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.Index)
	})
}
