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
	engine := crawlers.Register("1270", "PaulTan", "https://paultan.org/")

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        false,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)
	engine.OnLaunch(func() {
		baseURL := "https://paultan.org/topics/local-news/page/"
		for i := 0; true; i++ {
			if engine.Test != nil && engine.Test.Done {
				return
			}

			pageURL := baseURL + strconv.Itoa(i)
			resp, err := http.Get(pageURL)
			if err != nil {
				continue
			}

			dom, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				continue
			}
			urls := dom.Find(".xsmall > strong > a")
			if len(urls.Nodes) == 0 {
				break
			}
			urls.Each(func(i int, selection *goquery.Selection) {
				pageURL, ok := selection.Attr("href")
				if !ok {
					return
				}
				engine.Visit(pageURL, crawlers.News)
				engine.OnHTML(".post-content > p  img", func(element *colly.HTMLElement, ctx *crawlers.Context) {
					ctx.Image = append(ctx.Image, element.Attr("src"))
				})
			})

			err = resp.Body.Close()
			if err != nil {
				continue
			}
		}
	})

}
