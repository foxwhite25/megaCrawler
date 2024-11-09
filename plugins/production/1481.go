package production

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1481", "澳大利亚退伍军人事务部", "https://www.dva.gov.au/")

	engine.SetStartingURLs([]string{
		"https://www.dva.gov.au/sitemap.xml?page=1",
		"https://www.dva.gov.au/sitemap.xml?page=2",
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
		if strings.Contains(element.Text, "/news/") || strings.Contains(element.Text, "/media/") ||
			strings.Contains(element.Text, "/newsroom/") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("time.datetime", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(`.contextual-region.node > div > p, .contextual-region.node > div > ul,
	.contextual-region.node > div > div > p`,
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
