package storage

import (
	"bytes"
	"strings"
	"unicode"

	"github.com/gocolly/colly/v2"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

func removeControlChars(b []byte) []byte {
	var buf bytes.Buffer
	for _, r := range string(b) {
		if !unicode.IsControl(r) {
			buf.WriteRune(r)
		}
	}
	return buf.Bytes()
}

func init() {
	engine := crawlers.Register("1425", "The Nation (Thailand)", "https://www.nationthailand.com/")

	engine.SetStartingURLs([]string{
		"https://api.nationthailand.com/sitemap",
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

	// 存在非法xml
	extractorConfig.Apply(engine)

	engine.OnResponse(func(response *colly.Response, ctx *crawlers.Context) {
		contentType := strings.ToLower(response.Headers.Get("Content-Type"))
		isXMLFile := strings.HasSuffix(strings.ToLower(response.Request.URL.Path), ".xml") || strings.HasSuffix(strings.ToLower(response.Request.URL.Path), ".xml.gz")
		if !strings.Contains(contentType, "html") && (!strings.Contains(contentType, "xml") && !isXMLFile) {
			return
		}
		response.Body = removeControlChars(response.Body)
	})

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "sitemap") {
			engine.Visit(element.Text, crawlers.Index)
			return
		}
		engine.Visit(element.Text, crawlers.News)
	})

	engine.OnHTML(".section-1 > div > div > h2", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Text
	})

	engine.OnHTML(".detail > div > div > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
