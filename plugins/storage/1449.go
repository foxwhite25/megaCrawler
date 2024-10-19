package storage

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1449", "新印度快报", "https://www.newindianexpress.com/")

	engine.SetStartingURLs([]string{"https://www.newindianexpress.com/sitemap.xml"})

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
		if strings.Contains(element.Text, "sitemap-daily") {
			engine.Visit(element.Text, crawlers.Index)
		}

		keywords := []string{"nation", "world", "business", "states", "options",
			"cities", "sport", "good-news"} // 关键词列表

		for _, keyword := range keywords {
			if strings.Contains(element.Text, keyword) {
				engine.Visit(element.Text, crawlers.News)
				break // 找到匹配后跳出循环
			}
		}
	})

	engine.OnHTML("div.p-alt.arr--sub-headline.arrow-component.subheadline-m_subheadline__3fd7z.subheadline-m_dark__28u00",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Description += element.Text
		})

	engine.OnHTML(".arrow-component.arr--text-element.text-m_textElement__e3QEt.text-m_dark__1TC18 > p",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
