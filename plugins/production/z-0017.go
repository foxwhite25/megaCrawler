package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("z-0017", "medeka", "https://www.merdeka.com/")

	engine.SetStartingURLs([]string{
		"https://www.merdeka.com/peristiwa",
		"https://www.merdeka.com/trending",
		"https://www.merdeka.com/politik",
		"https://www.merdeka.com/uang",
		"https://www.merdeka.com/dunia",
		"https://www.merdeka.com/sepakbola",
		"https://www.merdeka.com/otomotif",
		"https://www.merdeka.com/artis"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)
	engine.OnHTML("span.item-title > a,a.mb-1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML("div.article > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	engine.OnHTML("time.text-xs > span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		team := strings.TrimSpace(element.Text)
		ctx.PublicationTime += team
	})
	engine.OnXML("/html/body/div[4]/main/section[2]/div/div[4]/div/div[1]/div/span/a/text()", func(element *colly.XMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
}
