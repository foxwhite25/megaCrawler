package dev

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2413", "Australian Associated Press", "https://www.aap.com.au/")

	engine.SetStartingURLs([]string{
		"https://www.aap.com.au/wp-sitemap-posts-post-1.xml",
		"https://www.aap.com.au/wp-sitemap-posts-post-2.xml",
		"https://www.aap.com.au/wp-sitemap-posts-post-3.xml",
		"https://www.aap.com.au/wp-sitemap-posts-post-4.xml",
		"https://www.aap.com.au/wp-sitemap-posts-post-5.xml",
		"https://www.aap.com.au/wp-sitemap-posts-post-6.xml",
		"https://www.aap.com.au/wp-sitemap-posts-post-7.xml",
		"https://www.aap.com.au/wp-sitemap-posts-post-8.xml",
		"https://www.aap.com.au/wp-sitemap-posts-post-9.xml",
		"https://www.aap.com.au/wp-sitemap-posts-post-10.xml",
		"https://www.aap.com.au/wp-sitemap-posts-post-11.xml",
		"https://www.aap.com.au/wp-sitemap-posts-post-12.xml",
		"https://www.aap.com.au/wp-sitemap-posts-post-13.xml",
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

	engine.OnHTML("div.col-12.col-lg-9.info > span:nth-child(2)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML("div.c-article__content.e-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
