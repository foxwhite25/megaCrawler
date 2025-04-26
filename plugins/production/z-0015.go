package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("z-0015", "印度尼西亚媒体报", "https://mediaindonesia.com/")

	engine.SetStartingURLs([]string{
		"https://mediaindonesia.com/politik-dan-hukum",
		"https://mediaindonesia.com/ekonomi",
		"https://mediaindonesia.com/megapolitan",
		"https://mediaindonesia.com/internasional",
		"https://mediaindonesia.com/humaniora",
		"https://mediaindonesia.com/olahraga",
		"https://mediaindonesia.com/sepak-bola",
		"https://mediaindonesia.com/nusantara"})

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
	engine.OnHTML("li > div.text  > h3>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML("div.pagination > a:nth-child(2)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})
	engine.OnHTML("span.datetime", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})
	engine.OnHTML("div.article > p ", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		element.DOM.Find("strong").Remove()
		ditecttext := element.DOM.Text()
		ctx.Content += ditecttext
	})
}
