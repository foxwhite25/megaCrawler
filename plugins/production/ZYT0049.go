package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("ZYT0049", "stuff", "https://stuff.co.za")

	engine.SetStartingURLs([]string{"https://stuff.co.za/sitemap_index.xml"})

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
		if strings.Contains(element.Text, "post-sitemap") {
			engine.Visit(element.Text, crawlers.Index)
		} else if !strings.Contains(element.Text, "xml") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("span + a[rel=\"author\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	engine.OnHTML(".date > .post-date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Attr("datetime")
	})

	engine.OnHTML(".sub-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description += element.Text
	})

	engine.OnHTML(".post-content > :not(div, h4)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
