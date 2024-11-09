package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1781", "卫生及社会关怀部", "https://www.gov.uk/government/organisations/department-of-health-and-social-care")

	engine.SetStartingURLs([]string{"https://www.gov.uk/search/news-and-communications?organisations[]=department-of-health-and-social-care&parent=department-of-health-and-social-care"})

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

	engine.OnHTML(".govuk-link.gem-print-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".govspeak", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(".govuk-link.govuk-pagination__link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.Index)
	})
}
