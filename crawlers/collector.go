package crawlers

import (
	"time"

	"github.com/gocolly/colly/v2"
)

type (
	HTMLCallback  func(element *colly.HTMLElement, ctx *Context)
	XMLCallback   func(element *colly.XMLElement, ctx *Context)
	CollyHTMLPair struct {
		callback colly.HTMLCallback
		selector string
	}
)

type XMLPair struct {
	callback XMLCallback
	selector string
}
type CollectorConstructor struct {
	parallelLimit    *int
	domainGlob       string
	timeout          time.Duration
	startingURLs     []string
	startHandler     func()
	robotTxt         string
	htmlHandlers     []CollyHTMLPair
	xmlHandlers      []XMLPair
	responseHandlers []func(response *colly.Response, ctx *Context)
	errorHandler     colly.ErrorCallback
	launchHandler    func()
}

func retryRequest(r *colly.Request, maxRetries int) int {
	retriesLeft := maxRetries
	if x, ok := r.Ctx.GetAny("retriesLeft").(int); ok {
		retriesLeft = x
	}
	if retriesLeft > 0 {
		r.Ctx.Put("retriesLeft", retriesLeft-1)
		_ = r.Retry()
	}
	return retriesLeft
}
