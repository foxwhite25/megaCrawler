package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("0052", "UNIVERSITY&COLLEGES Member of AMA Education System", "news.amaes.edu.ph")

	engine.SetStartingURLs([]string{
		"https://news.amaes.edu.ph/search?q=a"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        true,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML("h3>a:nth-child(2)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
	engine.OnHTML(".post-body>p:nth-child(1),[style=\"clear: both; text-align: left;\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fulltextwithauthor := element.Text
		textwithauthor := strings.Split(fulltextwithauthor, " By: ")
		if len(textwithauthor) > 1 {
			dateelement := strings.Split(textwithauthor[1], " , ")[0]
			ctx.PublicationTime = dateelement
		}
	})
	engine.OnHTML(".post-body>p:nth-child(1),[style=\"clear: both; text-align: left;\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fulltextwithdate := element.Text
		textwithdate := strings.Split(fulltextwithdate, " By: ")
		if len(textwithdate) > 1 {
			dateelement := strings.Split(textwithdate[1], " , ")[1]
			ctx.PublicationTime = dateelement
		}
	})
	engine.OnHTML(".post-body>p,.post-body>div>div,p>span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	engine.OnHTML("#blog-pager-older-link>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})
}
