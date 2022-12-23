package main

import "fmt"

func main() {

}

func plusOne(digits []int) []int {
	if digits[len(digits)-1] < 9 {
		digits[len(digits)-1] = digits[len(digits)-1] + 1

	} else {
		digits[len(digits)-1] = 0
		carry := 1
		i := len(digits) - 2
		for carry == 1 {
			if i < 0 && carry == 1 {

				digits = append([]int{1}, digits...)
				return digits
			}
			val := digits[i] + carry
			fmt.Println(i, ">>>", val)
			if val == 10 {
				digits[i] = 0
			} else {
				digits[i] = val
				return digits
			}
			// x = append([]int{1}, x...)
			i--
		}
	}
	return digits
}
