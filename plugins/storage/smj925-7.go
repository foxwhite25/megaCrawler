
package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("smj925-7", "Net-25", "http://www.feu-nrmf.edu.ph/")
	
	engine.SetStartingURLs([]string{
		"http://www.feu-nrmf.edu.ph/featured-stories",
		"http://www.feu-nrmf.edu.ph/newsletter"})
	
	extractorConfig := extractors.Config{
		Author:       true,//无作者
		Image:        true,//具体新闻页无时间，时间在外面不会采
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}
	
	extractorConfig.Apply(engine)
	engine.OnHTML(".title>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)
	})
	
	engine.OnHTML("a[rel=\"next\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.Index)
	})

	
	// 采集pdf文件
		engine.OnHTML(".views-field.views-field-nothing>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
    	fileURL := element.Attr("href")
    	if true {
        	url, err := element.Request.URL.Parse(fileURL)
        	if err != nil {
            	crawlers.Sugar.Error(err.Error())
            	return
        	}
        	ctx.File = append(ctx.File, url.String())
        	ctx.PageType = crawlers.Report
    	}
	})
}
