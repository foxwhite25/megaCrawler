package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("mwq-023", "When the Abuser Goes to Work", "https://www.abusergoestowork.com")

	engine.SetStartingURLs([]string{"https://www.abusergoestowork.com/sitemap.xml"})

	extractorConfig := extractors.Config{
		Author:       true,
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

	engine.OnHTML("div.entry-content > p,div.entry > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML("div.post-info > span.date.published.time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})

}
