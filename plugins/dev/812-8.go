
package dev

import (
	"strings"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
		
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("812-8", "马来西亚监狱局", "https://archive.data.gov.my/data/ms_MY/dataset/?q=jabatan+penjara+malaysia")
	//"http://www.prison.gov.my/"的所有信息都在这个网页
	engine.SetStartingURLs([]string{"https://archive.data.gov.my/data/ms_MY/dataset/?q=jabatan+penjara+malaysia"})
	
	extractorConfig := extractors.Config{
		Author:       false,
		Image:        true,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}
	
	extractorConfig.Apply(engine)

	engine.OnHTML(".dataset-content>h2>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	// 采集csv、XLS、XLSX文件
	engine.OnHTML(".resource-url-analytics", func(element *colly.HTMLElement, ctx *crawlers.Context) {
    	fileURL := element.Attr("href")
    	if strings.Contains(fileURL, ".csv") || 
       		strings.Contains(fileURL, ".xls") || 
       		strings.Contains(fileURL, ".xlsx") {
        	url, err := element.Request.URL.Parse(fileURL)
        	if err != nil {
            	crawlers.Sugar.Error(err.Error())
            	return
        	}
        	ctx.File = append(ctx.File, url.String())
        	ctx.PageType = crawlers.Report
    	}
	})

	engine.OnHTML("li.active+li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.Index)
	})
}
