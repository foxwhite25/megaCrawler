package dev

import (
	"encoding/xml"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

type URLSet struct { //解析带有命名空间XML的sitemap
	URLs []struct {
		Loc string `xml:"loc"`
	} `xml:"url"` //xml映射url
}

func init() {
	engine := crawlers.Register("N-0037", "Báo Đầu Tư", "https://baodautu.vn/")

	engine.SetStartingURLs([]string{"https://baodautu.vn/sitemap.xml"})

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
		if strings.Contains(element.Text, ".xml") {
			engine.Visit(element.Text, crawlers.Index)
		}
	})

	engine.OnResponse(func(response *colly.Response, ctx *crawlers.Context) {
		var urlset URLSet
		err := xml.Unmarshal(response.Body, &urlset)
		if err != nil {
			return
		}
		for _, url := range urlset.URLs {
			engine.Visit(strings.TrimSpace(url.Loc), crawlers.News)
		}
	})

	engine.OnHTML("div.mr-auto > a.author", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	engine.OnHTML("div.mr-auto > span.post-time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		re := regexp.MustCompile(`\d{2}/\d{2}/\d{4}`)
		matchs := re.FindStringSubmatch(element.Text)
		ctx.PublicationTime = strings.Join(matchs, "")
	})

	engine.OnHTML("div.sapo_detail", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description += element.Text
	})

	engine.OnHTML("#content_detail_news > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
