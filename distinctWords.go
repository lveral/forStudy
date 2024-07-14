package main

import (
	"fmt"
	"strings"
)

func countDistinctWords(messages []string) int {
	wordsMap := make(map[string]int)
	for _, msg := range messages {
		msgList := strings.Fields(msg)
		for _, word := range msgList {
			word = strings.ToLower(word)
			if _, ok := wordsMap[word]; !ok {
				wordsMap[word] = 0
			}
			wordsMap[word]++
		}
	}
	return len(wordsMap)
}

func main() {
	test := []string{""}
	fmt.Println(countDistinctWords(test))

}
