package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2450", "6abc Action News", "https://6abc.com/")

	engine.SetStartingURLs([]string{
		"https://6abc.com/sitemap/news/interim-2024.xml",
		"https://6abc.com/sitemap/news/2023.xml",
		"https://6abc.com/sitemap/news/2022.xml",
		"https://6abc.com/sitemap/news/2021.xml",
		"https://6abc.com/sitemap/news/2022.xml",
		"https://6abc.com/sitemap/news/2020.xml",
		"https://6abc.com/sitemap/news/2019.xml",
		"https://6abc.com/sitemap/news/2018.xml",
		"https://6abc.com/sitemap/news/2017.xml",
		"https://6abc.com/sitemap/news/2016.xml",
		"https://6abc.com/sitemap/news/2015.xml",
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

	engine.OnHTML("div.xvlfx.ZRifP.TKoO.eaKKC.EcdEg.bOdfO.qXhdi.NFNeu.UyHES > p",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
