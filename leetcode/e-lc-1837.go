package main

import (
	"math"
	"strconv"
)

func sumBase(n int, k int) int {
	i := strconv.FormatInt(int64(n), k)
	// fmt.Printf("type of a is %T\n", digit)
	digit, err := strconv.Atoi(i)
	if err != nil {
		// ... handle error
		panic(err)
	}
	sum := 0
	for digit != 0 {
		x := digit % 10
		sum += x
		digit = int(math.Floor(float64(digit) / 10.00))
		// fmt.Println(digit)
	}
	return sum
}
