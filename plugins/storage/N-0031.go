package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0031", "Quân đội Nhân dân", "https://www.qdnd.vn")

	engine.SetStartingURLs([]string{"https://www.qdnd.vn/sitemaps/pid/feed-0.xml"})

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
		if strings.Contains(element.Text, "/sitemaps/0/20") {
			engine.Visit(element.Text, crawlers.Index)
		} else if !strings.Contains(element.Text, ".xml") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("span.post-subinfo > span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		re := regexp.MustCompile(`\d{2}/\d{2}/\d{4}`) //时间正则表达式
		matchs := re.FindStringSubmatch(element.Text)
		ctx.PublicationTime = strings.Join(matchs, "")
	})

	engine.OnHTML("div.post-content > div > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
