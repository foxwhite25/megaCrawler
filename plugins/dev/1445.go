package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1445", "北京周报", "https://www.bjreview.com/")

	engine.SetStartingURLs([]string{
		"https://www.bjreview.com/China/",
		"https://www.bjreview.com/World/",
		"https://www.bjreview.com/Business/",
		"https://www.bjreview.com/Lifestyle/",
		"https://www.bjreview.com/Opinion/Governance/",
		"https://www.bjreview.com/Opinion/Pacific_Dialogue/",
		"https://www.bjreview.com/Opinion/Fact_Check/",
		"https://www.bjreview.com/Opinion/Voice/",
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

	//补全URL
	engine.OnHTML("a.bt-2020", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)
	})

	engine.OnHTML(`body > div > div > div > table > tbody > tr > td > a[target="_self"]`,
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			url, err := element.Request.URL.Parse(element.Attr("href"))
			if err != nil {
				crawlers.Sugar.Error(err.Error())
				return
			}
			engine.Visit(url.String(), crawlers.Index)
		})

	//匹配多种content样式
	engine.OnHTML(".TRS_Editor > p, .TRS_Editor > div > p, .TRS_Editor > div > div > span, .TRS_Editor > div > div > p, .TRS_Editor > div > div > div > p",
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})

}
