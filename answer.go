package main

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/McKael/madon"
)

var (
	randomAnswers = []string{
		"ㅇㅇ", "?", "응?", "응응", "응...", "뭐여?",
		"아니", "알긋다", "참말로 모르겠고만~",
		"그려", "괘안은데?", "왜 그려?", "등신이가", "치아뿌라 그거",
		"Seyana...", "그렇구나", "글쿠나", "글켔다",
	}

	seyana = []string{
		"Seyana...", "Soyana...", "Sorena...", "Arena...",
		"알긋다", "그려", "글켔다", "그러네...", "그랬지...", "그렇구나", "글쿠나",
	}

	why = []string{
		"뭐여?", "와 그러노?", "왜 그려?",
	}

	/*
		yesOrNo = []string{
			"응.", "ㅇㅇ", "아니.", "ㄴㄴ", "글쎼.", "등신이가",
		}
	*/
)

func akane(n *madon.Notification, content string) {
	var err error

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if strings.Contains(content, "아카네") {

		number := r.Intn(len(why))
		_, err = reply(n, why[number], "")

	} else if strings.Contains(content, ":seyana:") || strings.Contains(content, ":akane:") {
		number := r.Intn(len(seyana))
		_, err = reply(n, seyana[number], "")

		/*
			} else if strings.Contains(content, "어떻게") || strings.HasSuffix(content, "?") {
				_, err = reply(n, "참말로 모르겠고만~", "1")
		*/

	} else if strings.Contains(content, "가챠") || strings.Contains(content, "가샤") {
		_, err = reply(n, ":thinking_face:", "2")

	} else if strings.Contains(content, "할 거") || strings.Contains(content, "해 보") || strings.Contains(content, "생각") {
		number := r.Intn(len(seyana))
		_, err = reply(n, seyana[number], "")

	} else if strings.HasSuffix(content, "!") {
		number := r.Intn(len(randomAnswers))
		_, err = reply(n, randomAnswers[number]+"!", "")

	} else {

		number := r.Intn(len(randomAnswers))
		if yes := r.Intn(100); yes%3 == 0 {
			_, err = reply(n, randomAnswers[number], "")
		} else {
			_, err = reply(n, "응.", "")
		}

	}

	if err != nil {
		log.Println(err)
	}
}
