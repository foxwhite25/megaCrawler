
package production

import (
	"time"
	"strings"
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
					
	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("smj826-8", "菲尔生物", "https://www.philbio.org.ph/")
	
	engine.SetStartingURLs([]string{"https://www.philbio.org.ph/ourwork/"})
	
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
	engine.SetTimeout(60 * time.Second)


	engine.OnHTML(".entry-title>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		engine.Visit(url.String(), crawlers.News)
		time.Sleep(2* time.Second) 
	})

	engine.OnHTML("li.active+li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.Index)
	})

	engine.OnHTML(".author-name.fn>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})
}
