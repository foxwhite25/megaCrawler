package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("stanford", "斯坦福大学",
		"https://stanford.edu/")

	w.SetStartingURLs([]string{
		"https://news.stanford.edu/section/science-technology/",
		"https://news.stanford.edu/section/social-sciences/",
		"https://news.stanford.edu/section/law-policy/a",
		"https://www.gsb.stanford.edu/insights",
		"https://med.stanford.edu/news.html",
		"https://ed.stanford.edu/news-media",
	})

	w.OnHTML("html", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		extractors.Titles(ctx, element)
		extractors.Tags(ctx, element)
	})

	w.OnHTML(".su-news-components", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	partNews(w)
	partGsb(w)
	partMed(w)
	partEd(w)
}
