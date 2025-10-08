package dev

import (
	"encoding/json"
	"fmt"
	"io"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"net/http"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Response_A11 struct {
	Results []ResultItem_A11 `json:"results"`
}

type ResultItem_A11 struct {
	Alias string `json:"alias"`
}

func FetchAndVisitArticles_A11(engine *crawlers.WebsiteEngine, page int) {
	url := fmt.Sprintf("https://centurypacific.com.ph/wp-admin/admin-ajax.php", page)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var jsonResp Response_A11
	if err := json.Unmarshal(body, &jsonResp); err != nil {
		return
	}

	for _, item := range jsonResp.Results {
		fullURL := "https://www.centurypacific.com.ph" + item.Alias
		engine.Visit(fullURL, crawlers.News)
	}
}

func init() {
	engine := crawlers.Register("T-049", "Cebuana Lhuillier", "https://www.centurypacific.com.ph")

	engine.SetStartingURLs([]string{"https://centurypacific.com.ph/whats-new/"})

	extractorConfig := extractors.Config{
		Author:       false,
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
		for page := 14; page <= 40000; page += 10 {
			FetchAndVisitArticles_A11(engine, page)
		}
	})

	engine.OnHTML(".elementor-post__title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".elementor-element.elementor-widget-theme-post-content p > img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = []string{element.Attr("src")}
	})

	engine.OnHTML(".elementor-element.elementor-element-d7d9be6.elementor-align-left.elementor-widget.elementor-widget-post-info > div > ul > li > a > span > time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(".elementor-element.elementor-widget-theme-post-content p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

	engine.OnHTML(".elementor-widget__width-auto.elementor-widget.elementor-widget-text-editor > div > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})
}
