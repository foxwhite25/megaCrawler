package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0057", "Ministry of Defence (Malaysia)", "https://www.mod.gov.my/")

	engine.SetStartingURLs([]string{"https://www.mod.gov.my/index.php/en/media3/news"})

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

	engine.OnHTML(" td.list-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("li.page-item > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("div.article-details > div > p:nth-child(2)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		text := strings.TrimSpace(element.Text)
		re := regexp.MustCompile(`\b([1-9]|[12][0-9]|3[01]) (January|February|March|April|May|June|July|August|September|October|November|December) \d{4}\b`)
		match := re.FindString(text)
		ctx.PublicationTime = match
	})

	engine.OnHTML(`meta[name="author"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Attr("content"))
	})

	engine.OnHTML("div.article-details > div > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
