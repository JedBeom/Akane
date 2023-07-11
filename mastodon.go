package main

import (
	"fmt"
	"strings"

	madon "github.com/McKael/madon/v3"
	"github.com/microcosm-cc/bluemonday"
)

func filterContent(raw string) (content string) {
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

func reply(mc *madon.Client, status *madon.Status, reactions []Reaction) {
	content := filterContent(status.Content)
	answer := getAnswer(content, reactions)

	switch answer {
	case "A0001":
		username := status.Account.Username
		format := "%s, %s..."
		answer = fmt.Sprintf(format, username[0:1], username)
	}

	answer = putMention(answer, status)
	visibility := adjustVisibility(status.Visibility)

	cmdPost := madon.PostStatusParams{
		Text:       answer,
		InReplyTo:  status.ID,
		Visibility: visibility,
	}

	mc.PostStatus(cmdPost)
}

func putMention(content string, status *madon.Status) string {
	prefix := "@" + status.Account.Acct + " "
	for _, mention := range status.Mentions {
		if mention.Username == "akane" {
			continue
		}
		prefix += "@" + mention.Acct + " "
	}

	return prefix + content
}

func adjustVisibility(visibility string) string {
	if visibility == "public" {
		return "unlisted"
	}

	return visibility
}
