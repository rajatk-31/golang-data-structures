package main

import (
	"strings"
)

func main() {

}

func detectCapitalUse(word string) bool {
	if strings.ToUpper(word) == word {
		return true
	} else if strings.Title(strings.ToLower(word)) == word {
		return true
	} else if strings.ToLower(word) == word {
		return true
	}

	return false
}
