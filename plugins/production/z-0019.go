package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("z-0019", "BERITASATU", "https://www.beritasatu.com/")

	engine.SetStartingURLs([]string{
		"https://www.beritasatu.com/nasional/indeks",
		"https://www.beritasatu.com/indeks",
		"https://www.beritasatu.com/terpopuler/indeks",
		"https://www.beritasatu.com/terpopuler/indeks",
		"https://www.beritasatu.com/nusantara/indeks",
		"https://www.beritasatu.com/nusantara/dki-jakarta/indeks",
		"https://www.beritasatu.com/nusantara/bali/indeks",
		"https://www.beritasatu.com/nusantara/banten/indeks",
		"https://www.beritasatu.com/nusantara/jabar/indeks",
		"https://www.beritasatu.com/nusantara/jateng/indeks",
		"https://www.beritasatu.com/internasional/indeks",
		"https://www.beritasatu.com/sport/indeks",
		"https://www.beritasatu.com/lifestyle/indeks"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)
	engine.OnHTML("div.col-4 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML("li:has(a.active)+li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})
	engine.OnHTML("div.body-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	engine.OnHTML("body > main > div > div > div.col > small", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += strings.TrimSpace(element.Text)
	})
	engine.OnHTML("body > main > div > div > div.col > div.d-flex.my-3 > div.flex-grow-1.d-flex.pe-2 > div > b > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
}
