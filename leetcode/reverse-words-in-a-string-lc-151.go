package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseWords("Hello  world "))
}

func reverseWords(s string) string {
	// arr := strings.Split(s, " ")
	// final := ""
	// for i := len(arr) - 1; i >= 0; i-- {
	// 	if arr[i] != "" {
	// 		final = final + arr[i] + " "

	// 	}
	// }
	// return strings.TrimSpace(final)

	var Stack []string
	var temp = ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if len(temp) > 0 {
				Stack = append(Stack, temp)
			}
			temp = ""
		} else {
			temp = temp + string(s[i])
		}
	}
	if len(temp) > 0 {
		Stack = append(Stack, temp)
	}
	fmt.Println(Stack)
	var res string

	for k := len(Stack) - 1; k >= 0; k-- {
		res = res + Stack[k]
		if k == 0 {
			break
		}
		res = res + " "
	}
	return res
}

// func reverse(m *[]string, i int, j int) {
// 	for {
// 		if i > j {
// 			break
// 		}
// 		(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
// 		i++
// 		j--
// 	}
// }

// func reverseWords(s string) string {
// 	ss := strings.Fields(s)
// 	fmt.Println(ss, len(ss))
// 	reverse(&ss, 0, len(ss)-1)
// 	return strings.Join(ss, " ")
// }
