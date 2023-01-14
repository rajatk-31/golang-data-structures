package main

import "strings"

func toGoatLatin(sentence string) string {
	splitted := strings.Split(sentence, " ")
	s := "ma"
	for i, v := range splitted {
		s = s + "a"
		x := ""
		if v[0] == 'a' || v[0] == 'e' || v[0] == 'i' || v[0] == 'o' || v[0] == 'u' || v[0] == 'A' || v[0] == 'E' || v[0] == 'I' || v[0] == 'O' || v[0] == 'U' {
			x = v + s
		} else {
			x = v[1:] + v[:1] + s
		}
		splitted[i] = x

	}

	return strings.Join(splitted, " ")
}
