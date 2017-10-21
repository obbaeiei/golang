package golang

import "strings"

func LettersCount(sentense string) map[string]int {
	result := make(map[string]int)
	words := strings.Split(sentense, " ")
	for i := 0; i < len(words); i++ {
		word := words[i]
		result[word] = len(word)
	}
	return result
}
