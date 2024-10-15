package storage

import (
	"fmt"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1046", "台湾国立政治大学", "https://www.nccu.edu.tw/")

	engine.SetStartingURLs([]string{"https://www.nccu.edu.tw/p/403-1000-160-1.php?Lang=zh-tw"})

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

	engine.OnHTML(".more > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(fmt.Sprintf("https://www.nccu.edu.tw/p/403-1000-160-%s.php?Lang=zh-tw", strings.TrimSpace(element.Text)), crawlers.Index)
	})

	engine.OnHTML(".mtitle > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".meditor", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
