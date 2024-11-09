package production

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1478", "澳大利亚工业、创新与科学部", "https://www.industry.gov.au/")

	engine.SetStartingURLs([]string{"https://www.industry.gov.au/news"})

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

	engine.OnHTML("div.views-row > div > div > h3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)
	})

	engine.OnHTML(".pager__item--next > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})

	engine.OnHTML(".field.field--name-field-news-date.field--type-datetime.field--label-hidden.field__item",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		})

	engine.OnHTML("div.field.field--name-field-summary.field--type-string-long.field--label-hidden.field__item",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Description += element.Text
		})

	engine.OnHTML("div.clearfix.text-formatted.field.field--name-field-body.field--type-text-long.field--label-hidden.field__item",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
