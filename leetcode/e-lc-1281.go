package main

import "math"

func subtractProductAndSum(n int) int {
	sum := 0
	product := 1
	for n != 0 {
		x := n % 10
		sum += x
		product *= x
		n = int(math.Floor(float64(n) / 10.00))
		// fmt.Println(n)
	}
	return product - sum

}
