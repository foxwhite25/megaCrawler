package production

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1765", "股票代码报告", "https://www.tickerreport.com/")

	engine.SetStartingURLs([]string{"https://www.tickerreport.com/sitemap.xml"})

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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Request.URL.String(), "sitemap.xml") {
			if strings.Contains(element.Text, "post") {
				ctx.Visit(element.Text, crawlers.Index)
				return
			}
			return
		}
		ctx.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML(".entry", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
