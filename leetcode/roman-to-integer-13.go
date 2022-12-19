package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("LVIII"))
}

// func romanToInt(s string) int {
// 	roman := map[string]int{
// 		"I":    1,
// 		"II":   2,
// 		"III":  3,
// 		"IV":   4,
// 		"V":    5,
// 		"VI":   6,
// 		"VII":  7,
// 		"VIII": 8,
// 		"IX":   9,
// 		"X":    10,
// 		"XL":   40,
// 		"L":    50,
// 		"XC":   90,
// 		"C":    100,
// 		"CD":   400,
// 		"D":    500,
// 		"CM":   900,
// 		"M":    1000}
// 	last := 0
// 	no := 0
// 	data := strings.Split(s, "")
// 	for i := 0; i < len(data); i++ {
// 		if last < roman[data[i]] && i > 0 {
// 			no -= roman[data[i-1]]
// 			last = (roman[data[i]] - roman[data[i-1]])
// 			no += last
// 		} else {
// 			last = roman[data[i]]
// 			no += last
// 		}

// 		// if val, ok := roman["foo"]; ok {
// 		// 	//do something here
// 		// }
// 	}
// 	fmt.Println(roman)
// 	return no
// }

func romanToInt(s string) int {

	type Value struct {
		value  int
		before byte
	}

	dic := map[byte]Value{
		'I': {1, ' '}, 'V': {5, 'I'}, 'X': {10, 'I'}, 'L': {50, 'X'}, 'C': {100, 'X'}, 'D': {500, 'C'}, 'M': {1000, 'C'},
	}

	sum := 0
	before := byte(' ')
	length := len(s)

	for i := 0; i < length; i++ {

		d := dic[s[i]]
		sum += d.value

		if before == d.before {
			sum -= dic[d.before].value * 2
		}

		before = s[i]

	}
	return sum
}
