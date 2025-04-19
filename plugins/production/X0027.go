package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0027", "下议院领袖办公室", "https://www.gov.uk/government/organisations/the-office-of-the-leader-of-the-house-of-commons")
	// 网站http://www.commonsleader.gov.uk/变更为https://www.gov.uk/government/organisations/the-office-of-the-leader-of-the-house-of-commons
	engine.SetStartingURLs([]string{"https://www.gov.uk/search/news-and-communications?organisations[]=the-office-of-the-leader-of-the-house-of-commons&parent=the-office-of-the-leader-of-the-house-of-commons"})

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

	engine.OnHTML(".gem-c-document-list__item-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".govuk-grid-column-two-thirds.metadata-column > div > dl > dd:nth-child(4)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(".govspeak", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("a").Remove()
		directText := element.DOM.Text()
		ctx.Content += directText
	})

	engine.OnHTML(".govuk-link.govuk-pagination__link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.Index)
	})
}
