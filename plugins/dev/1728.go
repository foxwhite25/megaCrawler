package dev

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1728", "中东和北非金融网络", "https://menafn.com/")

	engine.SetStartingURLs([]string{
		"https://menafn.com/tb_news.aspx?tb_html=tb_us&tb_title=United%20States%20-%20Latest%20News",
		"https://menafn.com/tb_news.aspx?tb_html=tb_europe&tb_title=Europe%20-%20Latest%20News",
		"https://menafn.com/tb_news.aspx?tb_html=tb_arab&tb_title=Arab%20-%20Latest%20News",
		"https://menafn.com/tb_news.aspx?tb_html=tb_asia&tb_title=Asia%20-%20Latest%20News",
		"https://menafn.com/tb_news.aspx?tb_html=tb_politics&tb_title=Politics%20-%20Latest%20News",
	})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".entry-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	// engine.OnHTML("", func(element *colly.HTMLElement, ctx *crawlers.Context) {
	// 	ctx.Content += element.Text
	// })
	// 若采集到空文章，请将上述三行代码取消注释，并将Text的true改为false。

}
