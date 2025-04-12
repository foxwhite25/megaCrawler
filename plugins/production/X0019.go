package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("X0019", "Agency Tunis Afrique Press", "https://www.tap.info.tn/en/")

	engine.SetStartingURLs([]string{
		"https://www.tap.info.tn/en/portal%20-%20politics",
		"https://www.tap.info.tn/en/portal%20-%20economy",
		"https://www.tap.info.tn/en/portal_sciences_technology_eng",
		"https://www.tap.info.tn/en/portal%20-%20society",
		"https://www.tap.info.tn/en/portal%20-%20regions",
		"https://www.tap.info.tn/en/portal%20-%20world",
	})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        false,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".NewsItemText > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".NewsItemHeadline", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})

	engine.OnHTML(".HeaderSmall + tr .NewsItemText", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})
	// 查看所有内容需登录
	engine.OnHTML(".NewsItemCaption", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(".HeaderLarge + tr > td > table > tbody > tr:nth-last-child(1) > td > a[style='color:red;font-size:22px;'] + a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.Index)
	})
}
