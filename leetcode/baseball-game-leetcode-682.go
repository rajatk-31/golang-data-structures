package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
}

func calPoints(operations []string) int {
	var final []int
	for i := 0; i < len(operations); i++ {
		fmt.Println(final)
		inte, err := strconv.Atoi(operations[i])
		if err != nil {
			if operations[i] == "C" {
				if len(final) > 0 {
					final = final[:len(final)-1]
				}
			} else if operations[i] == "D" {
				if len(final) > 0 {
					final = append(final, final[len(final)-1]*2)
				}
			} else if operations[i] == "+" {
				if len(final) > 1 {
					final = append(final, final[len(final)-1]+final[len(final)-2])
				}
			} else if len(final) > 0 {
				final = append(final, final[len(final)-1])
			}

		} else {
			final = append(final, inte)
		}

	}
	var res int
	for j := 0; j < len(final); j++ {
		res = res + final[j]
	}
	return res
}
