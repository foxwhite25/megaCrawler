
package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
			
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("smj826-4", "RAFI拉蒙·阿博伊蒂兹基金会公司", "https://www.rafi.org.ph/")
	
	engine.SetStartingURLs([]string{"https://www.rafi.org.ph/?s=&years=&cat=rafinforms&tag=#uabb-search-results"})
	
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

	engine.OnHTML(".uabb-post-heading.uabb-blog-post-section>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML("div.pagination-next", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
