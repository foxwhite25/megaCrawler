package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("ZYT0056", "越南政府办公厅", "https://vpcp.chinhphu.vn/")
	
	engine.SetStartingURLs([]string{"https://vpcp.chinhphu.vn/tin-noi-bat.htm"})
	
	extractorConfig := extractors.Config{
		Author:       false, //no author
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}
	
	extractorConfig.Apply(engine)

	engine.OnHTML(".title-top>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".detail-content>p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(".detail-time>div", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})

}
