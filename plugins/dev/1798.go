package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1798", "司法部", "https://www.gov.uk/government/organisations/ministry-of-justice")

	// 司法部原网址：http://www.justice.gov.uk/，相关新闻指向第二、三个链接相关的部门。
	// 司法部现网址：https://www.gov.uk/government/organisations/ministry-of-justice，相关新闻指向第一个链接。

	engine.SetStartingURLs([]string{
		"https://www.gov.uk/search/news-and-communications?organisations[]=ministry-of-justice&parent=ministry-of-justice",
		"https://www.gov.uk/search/news-and-communications?organisations[]=hm-courts-and-tribunals-service&parent=hm-courts-and-tribunals-service",
		"https://www.gov.uk/search/news-and-communications?organisations[]=hm-prison-and-probation-service&parent=hm-prison-and-probation-service",
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

	engine.OnHTML("#js-results a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".govspeak", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(".govuk-link.govuk-pagination__link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.Index)
	})
}
