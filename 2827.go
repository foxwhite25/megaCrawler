package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2827", "GOV.PH", "https://web.nlp.gov.ph")

	engine.SetTimeout(60 * time.Second)

	engine.SetStartingURLs([]string{"https://web.nlp.gov.ph/category/events/"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(" div.uk-margin.uk-text-center > div > div > div > a ", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("#tm-main > div.uk-section-default.uk-section > div > div > div > div > ul > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("div > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		re := regexp.MustCompile(`\s+`)
		textWithSingleSpaces := re.ReplaceAllString(element.Text, " ")
		trimmedText := strings.TrimSpace(textWithSingleSpaces)
		ctx.Content += trimmedText
	})

}
