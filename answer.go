package main

import (
	"log"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"

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

	yesOrNo = []string{
		"응", "ㅇㅇ", "아니", "ㄴㄴ", "글쎄", "등신이가",
	}

	rate = []string{
		"괘안은데?", "이상하다", "별로인걸", "등신이가", "치아뿌라 그거", "참말로 모르겠고만~",
		"그런 건 잘 모르겠다는 걸~", "글켔다", "알긋다", "등신 아이가",
	}
)

func akane(n *madon.Notification, content string) {
	var err error

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 아카네짱!

	if strings.Contains(content, "민트") || strings.Contains(content, "민초") || strings.Contains(content, "초민") {

		if strings.Contains(content, "먹자") || strings.Contains(content, "아이스") ||
			strings.Contains(content, "먹을") || strings.Contains(content, "왔") {
			number := r.Intn(2)
			answer := ""
			switch number {
			case 0:
				answer = "다른 것은?"
			case 1:
				answer = n.Account.DisplayName + "...?"
			}

			_, err = reply(n, answer)
		} else {
			answer := []string{"민트라니 정신 나갔나", "그렇구나"}
			number := r.Intn(len(answer))
			_, err = reply(n, answer[number])
		}

	} else if strings.Contains(content, "아카네") &&
		utf8.RuneCountInString(content) < 6 && !strings.HasSuffix(content, "?") {

		number := r.Intn(len(why))
		_, err = reply(n, why[number])

		// :seyana: or :akane:
	} else if strings.Contains(content, ":seyana:") || strings.Contains(content, ":akane:") {
		number := r.Intn(len(seyana))
		_, err = reply(n, seyana[number])

		// 너 ~~~~지?
	} else if strings.Contains(content, "너") && strings.HasSuffix(content, "지?") {
		number := r.Intn(len(yesOrNo))
		_, err = reply(n, yesOrNo[number])

		// 할 거야
		// 해 보려고
		// 좋은 생각이
	} else if strings.Contains(content, "할") || strings.Contains(content, "해") ||
		strings.Contains(content, "생각") || strings.Contains(content, "어때") {

		number := r.Intn(len(rate))
		_, err = reply(n, rate[number])

	} else {

		number := r.Intn(len(randomAnswers))
		if yes := r.Intn(100); yes%3 == 0 {

			answer := randomAnswers[number]

			// ~~~!
			if strings.HasSuffix(content, "!") {
				answer += "!"
			}

			_, err = reply(n, answer)

		} else {
			_, err = reply(n, "응")
		}

	}

	if err != nil {
		log.Println(err)
	}
}
