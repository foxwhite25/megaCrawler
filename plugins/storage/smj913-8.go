
package dev

import (
	"strings"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("smj913-8", "3449", "http://www.mpic.com.ph/")
	
	engine.SetStartingURLs([]string{
		"https://www.mpic.com.ph/news/",
		"https://www.mpic.com.ph/investor-relations/annual-report/"})
	
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
	engine.OnHTML(".entry-title>a ", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	
	engine.OnHTML(".nextpostslink", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
	
	//采年报的pdf
	engine.OnHTML(".alink.customize-unpreviewable", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fileURL := element.Attr("href")
		if strings.Contains(fileURL, ".pdf") {
			url, err := element.Request.URL.Parse(element.Attr("href"))
			if err != nil {
				crawlers.Sugar.Error(err.Error())
				return
			}
			ctx.File = append(ctx.File, url.String())
			ctx.PageType = crawlers.Report
		}
	})
}
