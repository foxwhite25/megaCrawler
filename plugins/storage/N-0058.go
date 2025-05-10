package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0058", "Ministry of Ethnic and Religious Affairs", "http://ubdt.gov.vn/")

	engine.SetStartingURLs([]string{
		"http://ubdt.gov.vn/tin-tuc/tin-tuc-su-kien/thoi-su-chinh-tri.htm",
		"http://ubdt.gov.vn/tin-tuc/tin-tuc-su-kien/kinh-te-xa-hoi.htm",
		"http://ubdt.gov.vn/tin-tuc/tin-tuc-su-kien/y-te-giao-duc.htm",
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

	engine.OnHTML("div.col-md-12.col-xs-16 > div.content > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)
	})

	engine.OnHTML("div.pagination > ul > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))

		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})

	engine.OnHTML("#BodyContent_ctl00_ctl01_lblDate", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML("p.pull-right > span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		author := strings.Trim(element.Text, "()") // 去除中英文括号
		ctx.Authors = append(ctx.Authors, author)
	})

	engine.OnHTML(`#divNewsDetails > p:not([style*="font-size: 12px;"])`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
