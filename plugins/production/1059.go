package production

import (
	"strings"

	"github.com/gocolly/colly/v2"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

func init() {
	engine := crawlers.Register("1784", "加拿大政府门户网", "https://www.canada.ca/en.html")

	engine.SetStartingURLs([]string{
		"https://www.canada.ca/sitemap.xml",
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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if ctx.URL == "https://www.canada.ca/sitemap.xml" {
			if strings.Contains(element.Text, "/fr/") {
				return
			}
			ctx.Visit(element.Text, crawlers.Index)
			return
		}
		ctx.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML("body > main[property=\"mainContentOfPage\"] > div", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	engine.OnHTML(".gc-dwnld-lnk", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
		ctx.PageType = crawlers.Report
	})
}
