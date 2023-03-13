package extractors

import (
	"fmt"
	"regexp"
	"strings"

	"megaCrawler/crawlers"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

var digits = regexp.MustCompile(`\d`)
var byStatement = regexp.MustCompile(`[bB][yY][:\s]|[fF]rom[:\s]`)
var nameTokens = regexp.MustCompile(`[^\w'\-.]`)

func containsDigits(d string) bool {
	return digits.MatchString(d)
}

func parseByLine(searchStr string) (authors []string) {
	searchStr = byStatement.ReplaceAllString(searchStr, "")
	searchStr = strings.TrimSpace(searchStr)
	tokens := nameTokens.Split(searchStr, -1)
	for i, token := range tokens {
		tokens[i] = strings.TrimSpace(token)
	}

	var curname []string
	delimiters := []string{"and", ",", ""}

	for _, token := range tokens {
		if crawlers.Contain(delimiters, token) {
			if len(curname) > 0 {
				authors = append(authors, strings.Join(curname, " "))
				curname = []string{}
			}
		} else if !containsDigits(token) {
			curname = append(curname, token)
		}
	}

	if len(curname) >= 2 {
		authors = append(authors, strings.Join(curname, " "))
	}

	return
}

// getAuthors Takes a candidate line of html or text

func getAuthors(dom *colly.HTMLElement) (authors []string) {
	for _, attr := range []string{"name", "rel", "itemprop", "class", "id"} {
		for _, val := range []string{"author", "byline", "dc.creator", "byl"} {
			content := ""
			selection := dom.DOM.Find(fmt.Sprintf("*[%s=\"%s\"]", attr, val))
			tag := goquery.NodeName(selection)
			if tag == "meta" {
				if k, ok := selection.Attr("content"); ok {
					content = k
				}
			} else {
				content = selection.Text()
			}
			if len(content) > 0 {
				authors = append(authors, parseByLine(content)...)
			}
		}
	}
	return crawlers.Unique(authors)
}

func Authors(ctx *crawlers.Context, dom *colly.HTMLElement) {
	ctx.Authors = getAuthors(dom)
}
