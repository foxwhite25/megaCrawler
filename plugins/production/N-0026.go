package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0026", "Guard Online", "https://www.guardonline.com/")

	engine.SetStartingURLs([]string{
		"https://www.guardonline.com/tncms/sitemap/editorial.xml",
		"https://www.guardonline.com/tncms/sitemap/editorial.xml?year=2024",
		"https://www.guardonline.com/tncms/sitemap/editorial.xml?year=2023",
		"https://www.guardonline.com/tncms/sitemap/editorial.xml?year=2022",
		"https://www.guardonline.com/tncms/sitemap/editorial.xml?year=2021",
		"https://www.guardonline.com/tncms/sitemap/editorial.xml?year=2020",
	})

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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "/sitemap/editorial.xml?date=") {
			engine.Visit(element.Text, crawlers.Index)
		} else if !strings.Contains(element.Text, ".xml") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("ul.list-inline > li.hidden-print > time:nth-child(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Attr("datetime"))
	})

	engine.OnHTML("div.subscriber-preview > p, div.subscriber-only > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
