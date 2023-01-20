package main

func differenceOfSum(nums []int) int {
	x, y := 0, 0
	for _, n := range nums {
		x += n
		for n > 0 {
			y += n % 10
			n /= 10
		}
	}
	if x < y {
		x = -x
	}
	return x - y
}
