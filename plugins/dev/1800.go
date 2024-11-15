package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1800", "威尔斯事务大臣办公室", "https://www.gov.uk/government/organisations/office-of-the-secretary-of-state-for-wales")
	// 部门更名为“威尔斯办公室”，采集脚本见1801.go
	engine.SetStartingURLs([]string{"https://www.gov.uk/search/news-and-communications?organisations[]=office-of-the-secretary-of-state-for-wales&parent=office-of-the-secretary-of-state-for-wales"})

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
