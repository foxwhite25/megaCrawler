package dev

import (
	"strings" 
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("smj812-2", "马来西亚投资发展局", "http://www.mida.gov.my/")
	
	engine.SetStartingURLs([]string{//还有event网页：“https://www.mida.gov.my/media-and-events/events/”因为动态加载采集不了
		"https://www.mida.gov.my/media-and-events/announcement-media-release/",
		"https://www.mida.gov.my/media-and-events/news/",
		"https://www.mida.gov.my/report/"})
	
	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}
	
	extractorConfig.Apply(engine)

	engine.OnHTML(".formBtn", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})
		// 采集pdf文件
		engine.OnHTML(".resource-url-analytics", func(element *colly.HTMLElement, ctx *crawlers.Context) {
    	fileURL := element.Attr("href")
    	if strings.Contains(fileURL, ".pdf") {
        	url, err := element.Request.URL.Parse(fileURL)
        	if err != nil {
            	crawlers.Sugar.Error(err.Error())
            	return
        	}
        	ctx.File = append(ctx.File, url.String())
        	ctx.PageType = crawlers.Report
    	}
	})

	engine.OnHTML(".financity-single-article-content>p+p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		currentURL := element.Request.URL.String()
		// 仅当URL为"https://www.mida.gov.my/media-and-events/news/"时采集时间
		if strings.Contains(currentURL, "/media-and-events/news/") {
			ctx.PublicationTime += element.Text
		}
	})
}
