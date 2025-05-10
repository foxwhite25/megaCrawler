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

type Response struct {
	Data []ResultItem `json:"data"`
}

type ResultItem struct {
	Slug string `json:"slug"`
}

func FetchAndVisitArticles(engine *crawlers.WebsiteEngine, page int) {
	url := fmt.Sprintf("https://business-cambodia.com/cms/items/articles?limit=10&sort=-date_created&page=%d&filter[category][slug]=news&fields=title,%%20thumbnail,%%20date_created,user_created.first_name,%%20user_created.last_name,%%20user_created.avatar,slug,%%20category.slug,%%20views,%%20category.name", page)
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
		fullURL := "https://business-cambodia.com/articles/" + item.Slug
		engine.Visit(fullURL, crawlers.News)
	}
}

func init() {
	engine := crawlers.Register("N-0052", "Business Cambodia", "https://business-cambodia.com/")

	engine.SetStartingURLs([]string{"https://business-cambodia.com/categories/news"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  false,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnLaunch(func() {
		for page := 1; page <= 163; page += 1 {
			FetchAndVisitArticles(engine, page)
		}
	})

	engine.OnHTML("#title > div > div.flex.justify-center.text-gray-500 > span:nth-child(3)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML("#article_author > div.text-center > div.font-bold", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	engine.OnHTML("div.article_body > p,div.article_body > ul", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
