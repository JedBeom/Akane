package main

import (
	"log"
	"strings"

	"github.com/McKael/madon"
)

func akane(n *madon.Notification, content string) {
	var err error

	if strings.Contains(content, "났어") {
		_, err = reply(n, "뭐여?", "")
	} else if strings.Contains(content, "될 거") || strings.Contains(content, "될거") {
		_, err = reply(n, "글쿠나", "")
	} else if strings.Contains(content, "천재니까요") {
		_, err = reply(n, "응?", "")
	} else if strings.Contains(content, "생각이") {
		_, err = reply(n, "알긋다", "")
	} else if strings.Contains(content, "들어줘") {
		_, err = reply(n, "왜 그려?", "")
	} else if strings.Contains(content, "그렇게 생각하") {
		_, err = reply(n, "등신이가", "")
	} else if strings.Contains(content, "에에에") {
		_, err = reply(n, "참말로 모르겠다", "")
	} else {
		_, err = reply(n, "응", "")
	}

	if err != nil {
		log.Println(err)
	}
}
