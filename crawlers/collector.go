package crawlers

import (
	"math"
	"math/rand"
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

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func retryRequest(r *colly.Request, maxRetries int) (int, int) {
	retriesLeft := maxRetries
	if x, ok := r.Ctx.GetAny("retriesLeft").(int); ok {
		retriesLeft = x
	}

	if retriesLeft > 0 {
		// Exponential backoff
		millSecRetry := int(math.Pow(2, float64(maxRetries-retriesLeft+1)) * 1000)
		jitter := randRange(-500, 1500) * randRange(1, millSecRetry/1000)
		delay := time.Duration(millSecRetry+jitter) * time.Millisecond
		time.Sleep(delay)

		r.Ctx.Put("retriesLeft", retriesLeft-1)
		_ = r.Retry()
		return retriesLeft, millSecRetry + jitter
	}
	return retriesLeft, -1
}
