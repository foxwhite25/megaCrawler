package rand

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("rand", "兰德公司", "https://www.rand.org/").
		SetStartingUrls([]string{"https://www.rand.org/pubs.html", "https://www.rand.org/content/rand/blog/jcr:content/par/bloglist.ajax.0.html"}).
		SetTimeout(20 * time.Second)

	s.OnHTML(".list > li > .text", func(element *colly.HTMLElement) {
		t, err := time.Parse("Jan 2, 2006", element.ChildText(".date"))
		if err != nil {
			t = time.Now()
		}
		s.AddUrl(element.ChildAttr(".title > a", "href"), t)
	})

	s.OnHTML(".pagination > li > a", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})

	s.OnHTML(".product-page-abstract", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".body-text", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".product-main", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".authors", func(element *colly.HTMLElement) {
		megaCrawler.SetAuthor(element, element.ChildText("a"))
	})
}
