package main

import (
	"math/rand"
	"strings"
)

func getAnswer(content string, reactions []Reaction) string {
	for i := range reactions {
		if includes(content, reactions[i].Keywords) {
			return pick(reactions[i].Answers) // TODO: A001, A002
		}
	}

	defaultAnswers := []string{"응", "응.", "응", "응", "응", "응", "응", "응", "응", "응", "응", "응.", "응!", "웅", "응응", "ㅇ", ".", "맞나", "에나", "맞다 아이가"}
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
