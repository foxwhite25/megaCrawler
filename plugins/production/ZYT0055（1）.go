package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("ZYT0055", "科学技术部部", "https://most.mst.gov.vn")
	
	engine.SetStartingURLs([]string{"https://most.mst.gov.vn/en/Category/23/news--events.aspx"})
	
	extractorConfig := extractors.Config{
		Author:       false,//no author
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         false,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}
	
	extractorConfig.Apply(engine)

	engine.OnHTML(".H3_iTem_ListNews>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".Paging > span:nth-child(5)>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".Around_News_Content>div>p,#divArticleDescription2 > div", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("img").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".News_Time_Post", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})
}
