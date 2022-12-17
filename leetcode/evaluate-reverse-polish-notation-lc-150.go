package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

func evalRPN(tokens []string) int {
	final := []int{}
	for i, val := range tokens {
		//+-*/
		finalLen := len(final)
		value := 0
		if val == "+" {
			value = final[finalLen-1] + final[finalLen-2]
			final = final[:finalLen-2]
			final = append(final, value)
		} else if val == "-" {
			value = final[finalLen-2] - final[finalLen-1]
			final = final[:finalLen-2]
			final = append(final, value)
		} else if val == "*" {
			value = final[finalLen-2] * final[finalLen-1]
			final = final[:finalLen-2]
			final = append(final, value)
		} else if val == "/" {
			value = final[finalLen-2] / final[finalLen-1]
			final = final[:finalLen-2]
			final = append(final, value)
		} else {
			i, _ := strconv.Atoi(tokens[i])
			final = append(final, i)
		}
	}
	return final[0]
}
