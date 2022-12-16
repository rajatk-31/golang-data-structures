package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSumOfLengths([]int{3, 2, 2, 4, 3}, 3))
}

func minSumOfLengths(arr []int, target int) int {
	var dp [100001]int
	i := 0
	j := 0
	sum := 0
	minLen := 100001
	output := 100001

	for j < len(arr) {
		sum += arr[j]

		for sum > target {
			sum -= arr[i]
			i++
		}
		if sum == target {
			len := j - i + 1
			minLen = int(math.Min(float64(len), float64(minLen)))
			if i > 0 {
				output = int(math.Min(float64(output), float64(dp[i-1]+len)))
			}
		}
		dp[j] = minLen
		j++
	}
	if output == 100001 {
		output = -1
	}
	return output
}
