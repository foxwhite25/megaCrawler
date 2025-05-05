package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0050", "CamboJA News", "https://cambojanews.com")

	engine.SetTimeout(60 * time.Second)

	engine.SetStartingURLs([]string{
		"https://cambojanews.com/post-sitemap1.xml",
		"https://cambojanews.com/post-sitemap2.xml",
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

	engine.OnHTML("span.elementor-icon-list-text.elementor-post-info__item.elementor-post-info__item--type-custom > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	engine.OnHTML("div.elementor-element > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
