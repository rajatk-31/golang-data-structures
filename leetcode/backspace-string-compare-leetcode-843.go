package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(backspaceCompare("r##t", "#r#t"))
}

func backspaceCompare(s string, t string) bool {
	var w1 []rune
	var w2 []rune
	runes1 := []rune(s)
	for _, rune := range runes1 {
		// fmt.Println(i, rune)
		if rune == 35 {
			if len(w1) > 0 {
				w1 = w1[:len(w1)-1]
			}
		} else {
			w1 = append(w1, rune)
		}

	}
	runes2 := []rune(t)
	for _, rune2 := range runes2 {

		if rune2 == 35 {
			if len(w2) > 0 {
				w2 = w2[:len(w2)-1]
			}

		} else {
			w2 = append(w2, rune2)
		}

	}
	return reflect.DeepEqual(string(w1), string(w2))
}
