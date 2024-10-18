package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	//这网站的主域名为www.thenews.com.pk，它只提供当日的新闻，所以采取采集内部的一个专门提供新闻的模块
	engine := crawlers.Register("1450", "国际新闻", "https://www.thenews.com.pk/tns/")

	engine.SetStartingURLs([]string{
		"https://www.thenews.com.pk/tns/category/interviews",
		"https://www.thenews.com.pk/tns/category/dialogue",
		"https://www.thenews.com.pk/tns/category/special-report",
		"https://www.thenews.com.pk/tns/category/art-culture",
		"https://www.thenews.com.pk/tns/category/literati",
		"https://www.thenews.com.pk/tns/category/footloose",
		"https://www.thenews.com.pk/tns/category/political-economy",
		"https://www.thenews.com.pk/tns/category/sports",
		"https://www.thenews.com.pk/tns/category/shehr",
		"https://www.thenews.com.pk/tns/category/fashion",
		"https://www.thenews.com.pk/tns/category/encore",
		"https://www.thenews.com.pk/tns/category/instep",
		"https://www.thenews.com.pk/tns/category/in-the-picture",
	})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".w_c_left > div > ul > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(`.pagination_category > a[rel = "next"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".authorFullName > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	engine.OnHTML("div.detail-time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(".detail-desc > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})

}
