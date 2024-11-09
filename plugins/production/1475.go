package production

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1475", "澳大利亚财政部", "https://www.finance.gov.au/")

	engine.SetStartingURLs([]string{"https://www.finance.gov.au/about-us/news"})

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

	engine.OnHTML(".field-content > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)
	})

	engine.OnHTML(".page-item > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})

	engine.OnHTML(".file.file--mime-application-pdf.file--application-pdf > a, .node__content.clearfix > div > p > a",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
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

	engine.OnHTML(`.node__content.clearfix > div > p, .node__content.clearfix > div > ul, 
		.node__content.clearfix > div > div > p`,
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
