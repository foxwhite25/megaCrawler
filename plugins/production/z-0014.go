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
	engine := crawlers.Register("z-0014", "REPUBLIK", "https://www.republika.co.id/")

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
	engine.OnLaunch(func() {
		baseURL := "https://www.republika.co.id/ajax/latest_news/10/"
		for i := 20; true; i = i + 10 {
			if engine.Test != nil && engine.Test.Done {
				return
			}

			pageURL := baseURL + strconv.Itoa(i) + "/false/load_more//undefined"
			resp, err := http.Get(pageURL)
			if err != nil {
				continue
			}

			dom, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				continue
			}
			urls := dom.Find("li.list-group-item > a")
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

	engine.OnHTML("div.article-content p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
