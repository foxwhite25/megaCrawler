package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2462", "Above the Law", "https://abovethelaw.com/")

	engine.SetStartingURLs([]string{
		"https://abovethelaw.com/post-sitemap190.xml",
		"https://abovethelaw.com/post-sitemap191.xml",
		"https://abovethelaw.com/post-sitemap192.xml",
		"https://abovethelaw.com/post-sitemap193.xml",
		"https://abovethelaw.com/post-sitemap194.xml",
		"https://abovethelaw.com/post-sitemap195.xml",
		"https://abovethelaw.com/post-sitemap196.xml",
		"https://abovethelaw.com/post-sitemap197.xml",
		"https://abovethelaw.com/post-sitemap198.xml",
		"https://abovethelaw.com/post-sitemap199.xml",
		"https://abovethelaw.com/post-sitemap200.xml",
		"https://abovethelaw.com/post-sitemap201.xml",
		"https://abovethelaw.com/post-sitemap202.xml",
		"https://abovethelaw.com/post-sitemap203.xml",
		"https://abovethelaw.com/post-sitemap204.xml",
		"https://abovethelaw.com/post-sitemap205.xml",
		"https://abovethelaw.com/post-sitemap206.xml",
		"https://abovethelaw.com/post-sitemap207.xml",
		"https://abovethelaw.com/post-sitemap208.xml",
		"https://abovethelaw.com/post-sitemap209.xml",
		"https://abovethelaw.com/post-sitemap210.xml",
		"https://abovethelaw.com/post-sitemap211.xml",
		"https://abovethelaw.com/post-sitemap212.xml",
		"https://abovethelaw.com/post-sitemap213.xml",
	})

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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML("#single-post-content > div > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	engine.OnHTML("div.single-post__meta > div > a.author.url.fn", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	engine.OnHTML(".single-post__content > p, single-post__content > blockquote", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
