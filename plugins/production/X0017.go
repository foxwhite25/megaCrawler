package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0017", "新南威尔士大学", "https://www.unsw.edu.au/")

	engine.SetStartingURLs([]string{"https://www.unsw.edu.au/sitemap.xml"})

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
		if strings.Contains(element.Text, "www.unsw.edu.au/news") && strings.Contains(element.Text, "www.unsw.edu.au/newsroom/news") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".datetime", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(".aem-GridColumn--default--9 p,.aem-GridColumn--default--9 h2", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})
}
