package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2461", "ABC7 News", "https://abc7news.com/")

	engine.SetStartingURLs([]string{
		"https://abc7news.com/sitemap/news/interim-2024.xml",
		"https://abc7news.com/sitemap/news/2023.xml",
		"https://abc7news.com/sitemap/news/2022.xml",
		"https://abc7news.com/sitemap/news/2021.xml",
		"https://abc7news.com/sitemap/news/2020.xml",
		"https://abc7news.com/sitemap/news/2019.xml",
		"https://abc7news.com/sitemap/news/2018.xml",
		"https://abc7news.com/sitemap/news/2017.xml",
		"https://abc7news.com/sitemap/news/2016.xml",
		"https://abc7news.com/sitemap/news/2015.xml",
		"https://abc7news.com/sitemap/news/2014.xml",
		"https://abc7news.com/sitemap/news/2013.xml",
	})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML("figure.kzIjN.GNmeK.pYrtp.dSqFO > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	engine.OnHTML(".xvlfx.ZRifP.TKoO.eaKKC.EcdEg.bOdfO.qXhdi.NFNeu.UyHES > p, div.body-text > p",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
