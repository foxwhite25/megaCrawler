
package production

import (
	"strings"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("smj826-11", "2819", "https://www.greenpeace.org/philippines/")
	
	engine.SetStartingURLs([]string{"https://www.greenpeace.org/philippines/news-stories/?post-type=press"})
	
	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}
	
	extractorConfig.Apply(engine)

	engine.OnHTML(".query-list-item-headline.wp-block-post-title>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".wp-block-query-pagination-next", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".single-post-author", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	engine.OnHTML(".single-post-time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})
}
