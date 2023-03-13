package errors

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

func init() {
	engine := crawlers.Register("1140", "人权研究所", "https://web.law.columbia.edu")

	engine.SetStartingURLs([]string{})

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
}
