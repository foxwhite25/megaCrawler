package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2570", "Air Cargo World", "https://cargofacts.com/")

	engine.SetStartingURLs([]string{"https://cargofacts.com/sitemap_index.xml"})

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
		if strings.Contains(element.Text, "/post-sitemap") {
			engine.Visit(element.Text, crawlers.Index)
		} else if !strings.Contains(element.Text, ".xml") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("div.content-inner > figure > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	engine.OnHTML("div.content-inner > p:not(:has(b>i)), div.content-inner > ul",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			// 过滤noscript标签内容,避免可能的垃圾信息
			directText := element.DOM.Contents().Not("noscript").Text()
			ctx.Content += directText
		})
}
