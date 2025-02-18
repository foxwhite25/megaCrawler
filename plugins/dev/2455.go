package dev

import (
	"encoding/json"
	"log"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("2455", "AARP: The Magazine", "https://www.aarp.org/")

	engine.SetStartingURLs([]string{"https://www.aarp.org/sitemap.xml"})

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

	engine.OnHTML("div.aarpe-text-image > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	// 处理 <script type="application/ld+json"> 标签
	engine.OnHTML("script[type='application/ld+json']", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		var jsonData map[string]interface{}
		err := json.Unmarshal([]byte(element.Text), &jsonData)
		if err != nil {
			log.Println("Error parsing JSON-LD:", err)
			return
		}

		// 提取文本内容
		if picture, found := jsonData["image"].(string); found {
			ctx.Image = []string{picture}
		}
		if articleBody, found := jsonData["articleBody"].(string); found {
			ctx.Content += articleBody
		}
	})

}
