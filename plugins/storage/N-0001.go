package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0001", "Chinese Commercial News", "https://www.shangbao.com.ph/")

	engine.SetStartingURLs([]string{"https://www.shangbao.com.ph/"})

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

	engine.OnHTML(".conlist > .contt > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("#nav_left > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("div.left_time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		text := strings.TrimSpace(element.Text)
		re := regexp.MustCompile(`^\d{4}年\d{2}月\d{2}日`)
		if match := re.FindString(text); match != "" {
			ctx.PublicationTime = match
		}
	})

	engine.OnHTML("div.left_zw > #fontzoom > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
