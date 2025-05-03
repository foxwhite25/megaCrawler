package dev

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"net/http"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Article struct {
	Link string `json:"link"`
}

type Response struct {
	Data []Article `json:"data"`
}

// 等待增加新闻数据数量
func FetchAndVisitArticles(engine *crawlers.WebsiteEngine, page int) {
	url := fmt.Sprintf("https://api.posttoday.com/api/v1.0/categories/politics?page=%d", page)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("读取响应失败: %v\n", err)
		return
	}

	var jsonResp Response
	if err := json.Unmarshal(body, &jsonResp); err != nil {
		log.Printf("解析 JSON 失败: %v\n", err)
		return
	}

	for _, item := range jsonResp.Data {
		fullURL := item.Link
		if !strings.HasPrefix(fullURL, "http") {
			fullURL = "https://www.posttoday.com" + fullURL
		}
		engine.Visit(fullURL, crawlers.News)
	}
}

// AJAX请求采集
func init() {
	engine := crawlers.Register("N-0043", "Post Today", "https://www.posttoday.com/")

	engine.SetParallelism(2)

	engine.SetStartingURLs([]string{"https://www.posttoday.com/category/politics"})

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

	engine.OnLaunch(func() {
		for page := 1; page <= 103; page++ { //设置最大页数量
			FetchAndVisitArticles(engine, page)
		}
	})

	engine.OnHTML("div.content-detail > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.News {
			if !strings.Contains(element.Text, "English version") {
				ctx.Content += element.Text
			}
		}
	})
}
