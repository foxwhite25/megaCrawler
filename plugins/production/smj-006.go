package production

import (
	"regexp"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("smj-006", "KONICA_MINOLTA", "https://www.konicaminolta.com.cn/")

	engine.SetStartingURLs([]string{"https://www.konicaminolta.com.cn/aboutus/news/index.html"})

	extractorConfig := extractors.Config{
		Author:       false,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         false,
		Title:        false,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	re := regexp.MustCompile(`window\.open\('([^']+)'`)

	engine.OnHTML(".table.table-hover.news.list > tbody > tr >td:nth-of-type(1)> a ", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		onclick := element.Attr("onclick")
		matches := re.FindStringSubmatch(onclick)
		if len(matches) >= 2 {
			relativePath := matches[1]
			baseURL := element.Request.URL.Scheme + "://" + element.Request.URL.Host
			absoluteURL := baseURL + relativePath

			engine.Visit(absoluteURL, crawlers.News)
		}
	})

	engine.OnHTML(".row.newstext.container > h1:nth-of-type(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})

	engine.OnHTML(".row.newstext.container > h1:nth-of-type(2)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle += element.Text
	})

	engine.OnHTML(".newsfulltext > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
