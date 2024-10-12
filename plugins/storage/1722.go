package storage

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1722", "胡佛研究所", "https://www.hoover.org/")

	engine.SetStartingURLs([]string{"https://www.hoover.org/sitemap.xml"})

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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "news") {
			engine.Visit(element.Text, crawlers.News)
			return
		}
		if engine.VisitIfContains(element.Text, []string{"research", "publications", "commentary", "press-releases"}, crawlers.Report) {
			return
		}
	})

	engine.OnHTML(".col-three > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	// engine.OnHTML(".content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	ctx.Content += element.Text
	// })
	// 若采集到空文章，请将上述三行代码取消注释，并将Text的true改为false。

	engine.OnHTML("ais-Pagination-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
