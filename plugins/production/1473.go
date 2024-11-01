package production

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1473", "澳大利亚就业部", "https://www.dewr.gov.au/newsroom")

	engine.SetStartingURLs([]string{"https://www.dewr.gov.au/sitemap.xml?page=1"})

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

	// 这个网站和澳大利亚教育与培训部非常相似
	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "/announcements/") ||
			strings.Contains(element.Text, "/newsroom/articles/") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	// 公告和新闻不同的selector
	engine.OnHTML(".node__content.clearfix > div > time, .container > time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	engine.OnHTML(".node__content.clearfix > header > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description += element.Text
	})

	engine.OnHTML(`.paragraph.paragraph--type--text.paragraph--view-mode--default > div > p,
	.paragraph.paragraph--type--text.paragraph--view-mode--default > div > ul`,
		func(element *colly.HTMLElement, ctx *crawlers.Context) {
			ctx.Content += element.Text
		})
}
