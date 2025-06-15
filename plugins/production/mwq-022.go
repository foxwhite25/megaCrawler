package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("mwq-022", "FLYING magazine", "https://www.flyingmag.com")

	engine.SetStartingURLs([]string{"https://www.flyingmag.com/sitemap_index.xml"})

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

	extractorConfig.Apply(engine)
	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "post-sitemap") {
			engine.Visit(element.Text, crawlers.Index)
		} else if !strings.Contains(element.Text, ".xml") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("#entry-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(" div.entry-meta.flex.flex-wrap.text-12.leading-30.mb-20 > div > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
}
