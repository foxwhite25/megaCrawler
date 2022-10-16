package rand

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"strconv"
)

func init() {
	w := megaCrawler.Register("rand", "foo", "https://www.rand.org/")

	w.SetStartingUrls([]string{"https://www.rand.org/about/people.html", "https://www.rand.org/pubs.html"})

	w.OnHTML("#Pagination > ul > li:nth-child(13) > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		page, _ := strconv.Atoi(element.Text)
		for i := 2; i < page; i++ {
			u := fmt.Sprintf("https://www.rand.org/about/people.html?page=%d", i)
			w.Visit(u, megaCrawler.Index)
		}
	})

	w.OnHTML("#staff-ul > li > div > h3 > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Expert)
	})

	w.OnHTML("#RANDTitleHeadingId", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Name = element.Text
	})

	w.OnHTML("#srch > div.bio-meta.full-bg-gray > div > div.title", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML("#onebio_overview > div > div.biography", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Description = element.Text
	})

	w.OnHTML("#results > ul > li > div.text > h3 > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Report)
	})

	w.OnHTML("#results > div.pagination-wrap > ul > li > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})
}
