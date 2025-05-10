package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

// 存在双域名mic.gov.vn，mst.gov.vn
func init() {
	engine := crawlers.Register("N-0056", "Ministry of Science and Technology", "https://mst.gov.vn/")

	engine.SetStartingURLs([]string{"https://mst.gov.vn/sitemap.xml"})

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
		if strings.Contains(element.Text, "/sitemaps/sitemaps") {
			engine.Visit(element.Text, crawlers.Index)
		} else if !strings.Contains(element.Text, ".xml") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("div.detail-time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		text := strings.TrimSpace(element.Text)
		re := regexp.MustCompile(`(0[1-9]|[12][0-9]|3[01])/(0[1-9]|1[0-2])/\d{4}`)
		match := re.FindString(text)
		ctx.PublicationTime = match
	})

	engine.OnHTML(`meta[name="author"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Attr("content"))
	})

	engine.OnHTML("div.detail-content.afcbc-body > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
