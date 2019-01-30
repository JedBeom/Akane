package main

import (
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

func contentExtraction(raw string) (content string) {
	// Remove HTML tags
	contentRaw := bluemonday.StrictPolicy().Sanitize(raw)

	contentArray := strings.Split(contentRaw, " ")

	// Remove @...
	for _, value := range contentArray {
		if len(value) == 0 {
			continue
		}

		if string(value[0]) == "@" {
			continue
		}

		content += value + " "
	}

	if len(content) != 0 {
		content = content[:len(content)-1]
	}
	return
}
