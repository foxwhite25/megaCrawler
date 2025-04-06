package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1284", "SANEF", "https://sanef.org.za/")

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

	engine.OnLaunch(func() {
		baseURL := "https://sanef.org.za/category/news/page/"
		for i := 1; true; i++ {
			if engine.Test != nil && engine.Test.Done {
				return
			}

			pageURL := baseURL + strconv.Itoa(i) + "/?tf-scroll=1"
			resp, err := http.Get(pageURL)
			if err != nil {
				continue
			}

			dom, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				continue
			}
			urls := dom.Find(".post-image > a")
			if len(urls.Nodes) == 0 {
				break
			}
			urls.Each(func(i int, selection *goquery.Selection) {
				pageURL, ok := selection.Attr("href")
				if !ok {
					return
				}
				engine.Visit(pageURL, crawlers.News)

			})

			err = resp.Body.Close()
			if err != nil {
				continue
			}
		}
	})
	engine.OnHTML(".entry-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
