package main

import (
	"math/rand"
	"strings"
)

func getAnswer(content string, reactions []Reaction) string {
	if rand.Int()%4 == 0 { // 25%의 확률로 무조건 '응.'
		return "응."
	}

	for i := range reactions {
		if includes(content, reactions[i].Keywords) {
			return pick(reactions[i].Answers)
		}
	}

	defaultAnswers := []string{"응", "응", "응", "응", "응", "응!", "웅", "응응", "ㅇ", ".", "맞나", "에나", "맞다 아이가"}
	return pick(defaultAnswers)
}

func pick(strs []string) string {
	randomIndex := rand.Int() % len(strs)
	return strs[randomIndex]
}

func includes(text string, keywords []string) bool {
	for _, keyword := range keywords {
		if strings.Contains(text, keyword) {
			return true
		}
	}
	return false
}
