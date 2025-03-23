package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1244", "亚太时事杂志", "https://thediplomat.com/")

	engine.SetStartingURLs([]string{
		"https://thediplomat.com/topics/diplomacy/",
		"https://thediplomat.com/topics/economy/",
		"https://thediplomat.com/topics/environment/",
		"https://thediplomat.com/topics/opinion/",
		"https://thediplomat.com/topics/politics/",
		"https://thediplomat.com/topics/security/",
		"https://thediplomat.com/topics/society/"})

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
	engine.OnHTML(".td-posts > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".td-list-pager--bottom > div.td-list-pager__nav > a:nth-child(3)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
