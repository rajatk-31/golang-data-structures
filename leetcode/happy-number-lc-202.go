package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}

func isHappy(n int) bool {
	previous := map[int]bool{}
	check := n
	for check != 1 {
		fmt.Println("--", check)
		check = calculate(check)
		if _, ok := previous[check]; ok {
			return false
		}
		previous[check] = true
	}
	return true

}

func calculate(no int) int {
	nos := strings.Split(strconv.Itoa(no), "")
	sum := 0
	for i := 0; i < len(nos); i++ {
		val, _ := strconv.Atoi(nos[i])
		sum += (val * val)
	}

	return sum

}
