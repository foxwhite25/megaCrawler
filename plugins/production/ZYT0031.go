package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("ZYT0031", "mitsuichemicals", "https://ap.mitsuichemicals.com")

	engine.SetStartingURLs([]string{"https://ap.mitsuichemicals.com/sitemap.xml"})

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
		if strings.Contains(element.Text, "mc") && strings.Contains(element.Text, "_") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".news-details__publisher", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	engine.OnHTML("p.news-details__date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})

	engine.OnHTML(".mod-sideContent-inner > :not(img)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
