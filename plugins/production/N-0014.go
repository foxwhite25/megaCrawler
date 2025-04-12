package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("N-0014", "国际新闻", "https://www.thenews.com.pk/")

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
		Image:        false,
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
		cleanText := strings.ReplaceAll(element.Text, "\n", "") //清除换行符
		ctx.Content += cleanText
	})
}
