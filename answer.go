package main

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/McKael/madon"
)

var (
	answers = []string{
		"ㅇㅇ", "?", "응?", "응응", "뭐여?", "그렇구나", "글쿠나",
		"Seyana...", "Soyana...", "Sorena...", "Arena...",
		"알긋다", "그려", "글켔다", "아니",
		// "알긋다", "알긋다(감명)", "알긋다(예리)", "알긋다(득도)", "알긋다(박식)", "알긋다(천하무쌍)",
		"그려", "괘안은데?", "왜 그려?", "등신이가", "치아뿌라 그거",
	}
)

func akane(n *madon.Notification, content string) {
	var err error

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	number := r.Intn(len(answers))

	if strings.Contains(content, "생각이") {
		_, err = reply(n, "알긋다", "")
	} else if strings.Contains(content, ":seyana:") || strings.Contains(content, ":akane:") {
		_, err = reply(n, "SEYANA", "")
	} else if strings.Contains(content, "어떻게") {
		_, err = reply(n, "참말로 모르겠고만~", "")
	} else if strings.Contains(content, "가챠") || strings.Contains(content, "가샤") {
		_, err = reply(n, "일단 해보는 거여", "")
	} else if strings.HasSuffix(content, "!") {
		_, err = reply(n, answers[number]+"!", "")
	} else {

		if yes := r.Intn(100); yes%3 == 0 {
			_, err = reply(n, answers[number], "")
			return
		}

		_, err = reply(n, "응", "")

	}

	if err != nil {
		log.Println(err)
	}
}
