package storage

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1454", "ReleaseWire", "http://www.releasewire.com/")

	engine.SetTimeout(60 * time.Second) //延长等待等待服务器响应的时间

	engine.SetStartingURLs([]string{"http://www.releasewire.com/press-releases/sitemap.php"})

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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, ".php") {
			engine.Visit(element.Text, crawlers.Index)
		}
		if strings.Contains(element.Text, ".htm") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML("#prbody > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
