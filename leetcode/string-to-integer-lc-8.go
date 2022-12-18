package main

import "math"

func main() {

}

func myAtoi(s string) int {
	start := false
	flag := 1
	res := 0
	for _, v := range s {
		if v == ' ' && !start {
			continue
		}
		if v == '-' && !start {
			flag = -1
			start = true
			continue
		}
		if v == '+' && !start {
			start = true
			continue
		}

		if v >= 48 && v <= 57 {
			start = true
			digit := int(v-'0') * flag
			if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && digit > 7) {
				return math.MaxInt32
			}
			if res < math.MinInt32/10 || (res == math.MinInt32/10 && digit < -8) {
				return math.MinInt32
			}
			res = res*10 + digit
		} else {
			break
		}
	}
	return res
}
