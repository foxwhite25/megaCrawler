
package dev

import (
	"strings"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("smj913-3", "3437", "https://nlex.com.ph/")
	
	engine.SetStartingURLs([]string{"https://nlex.com.ph/news-and-events/"})
	
	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}
	
	extractorConfig.Apply(engine)


	engine.OnHTML(".entry-title>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)
	})

	engine.OnHTML(".pagination-next", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML("div.fusion-meta-info-wrapper>span:nth-of-type(3)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})

	engine.OnHTML(".fn>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

}
