
package dev

import (
	"strings"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("SR0037", "Usep", "https://www.usep.edu.ph/")

	engine.SetStartingURLs([]string{"https://www.usep.edu.ph/headlines-2/"})

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

	engine.OnHTML(" div.bdp-post-content > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("div > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML("a.next.page-numbers", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

    engine.OnHTML("div#posted_on_by > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fullText := strings.TrimSpace(element.Text)
		parts := strings.Split(fullText, "Posted by: ")
		if len(parts) > 1 {
			author := strings.TrimSpace(parts[1])
			ctx.Authors = append(ctx.Authors, author)
		}
	})
}