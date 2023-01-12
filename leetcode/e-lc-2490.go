package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isCircularSentence("Leetcode exercise sound delightful"))
}

func isCircularSentence(sentence string) bool {
	splitted := strings.Split(sentence, " ")
	splittedlen := len(splitted)
	for i := 0; i < splittedlen; i++ {

		currlen := len(splitted[i])
		current := splitted[i]
		var forward string
		if i == splittedlen-1 {
			forward = splitted[0]
		} else {
			forward = splitted[i+1]
		}
		if current[currlen-1] != forward[0] {
			return false
		}
	}

	return true
}
