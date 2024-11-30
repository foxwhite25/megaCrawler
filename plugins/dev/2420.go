package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2420", "European Union News", "https://european-union.europa.eu/")

	engine.SetStartingURLs([]string{
		"https://european-union.europa.eu/news-and-events/news-and-stories_en",
	})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("div.ecl-content-block__title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), "/news/") || strings.Contains(element.Attr("href"), "/press/") {
			engine.Visit(element.Attr("href"), crawlers.News)
		}
	})

	engine.OnHTML("li.ecl-pagination__item.ecl-pagination__item--next > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})

	//这个网站新闻页不同的网页结构非常多
	engine.OnHTML(`div.ecl > p, div.ecl > ul, 
		div.gsc-bge-grid__area > p,
		div.gsc-bge-grid__area > h4, div.gsc-bge-grid__area > ul:not([class]),
		div.card-body > div.clearfix.text-formatted > p,
		div.ep-a_text > p, 
		div.node__content.clearfix > div > p,
		div.section > p`,
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
