package dev

import (
	"fmt"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0044", "Khaosan Pathet Lao", "https://kpl.gov.la/")

	engine.SetTimeout(60 * time.Second)
	engine.SetStartingURLs([]string{"https://kpl.gov.la/news.aspx?cat=1&page=1"})

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

	engine.OnResponse(func(response *colly.Response, ctx *crawlers.Context) {
		// 只对索引页执行翻页
		if ctx.PageType != crawlers.Index {
			return
		}

		// 获取当前URL的page参数
		currentURL := response.Request.URL
		pageParam := currentURL.Query().Get("page")

		// 尝试解析当前page值
		pageValue, err := strconv.Atoi(pageParam)
		if err == nil && pageValue < 1314 {
			nextPageURL := fmt.Sprintf("https://kpl.gov.la/News.aspx?cat=1&page=%d", pageValue+1)
			engine.Visit(nextPageURL, crawlers.Index)
		}

	})

	engine.OnHTML("ul.news-story > li > h3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))

		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)
	})

	engine.OnHTML("span.post-time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML("div.post-ct-entry > p,div.post-ct-entry > div,div.post-ct-entry > span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
