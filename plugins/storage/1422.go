package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1422", "人民评论", "https://www.peoplesreview.com.np/")

	engine.SetStartingURLs([]string{
		"https://www.peoplesreview.com.np/wp-sitemap-posts-post-1.xml",
		"https://www.peoplesreview.com.np/wp-sitemap-posts-post-2.xml",
		"https://www.peoplesreview.com.np/wp-sitemap-posts-post-3.xml",
		"https://www.peoplesreview.com.np/wp-sitemap-posts-post-4.xml",
		"https://www.peoplesreview.com.np/wp-sitemap-posts-post-5.xml",
		"https://www.peoplesreview.com.np/wp-sitemap-posts-post-6.xml",
		"https://www.peoplesreview.com.np/wp-sitemap-posts-post-7.xml",
		"https://www.peoplesreview.com.np/wp-sitemap-posts-post-8.xml",
		"https://www.peoplesreview.com.np/wp-sitemap-posts-post-9.xml",
		"https://www.peoplesreview.com.np/wp-sitemap-posts-post-10.xml",
		"https://www.peoplesreview.com.np/wp-sitemap-posts-post-11.xml",
		"https://www.peoplesreview.com.np/wp-sitemap-posts-post-12.xml",
		"https://www.peoplesreview.com.np/wp-sitemap-posts-post-13.xml",
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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML(".uk-text-large.uk-text-justify > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
