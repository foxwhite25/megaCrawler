package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("ZYT0026", "pharmaphorum", "https://pharmaphorum.com")

	engine.SetStartingURLs([]string{"https://pharmaphorum.com/sitemap.xml"})

	extractorConfig := extractors.Config{
		Author:       false, //无作者
		Image:        false,
		Language:     true,
		PublishDate:  false, //无时间
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "sitemap") {
			engine.Visit(element.Text, crawlers.Index)
		} else if !strings.Contains(element.Text, "xml") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".entry-content p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
