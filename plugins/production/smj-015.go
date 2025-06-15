package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("smj-015", "Danske_Bank", "https://danskebank.com/")

	engine.SetStartingURLs([]string{"https://danskebank.com/sitemap.xml"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML("div.article-meta", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	engine.OnHTML(".article-content,p.article-subheadline", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
