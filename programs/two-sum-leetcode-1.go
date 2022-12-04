package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)

	for i, num := range nums {
		if _, ok := mp[target-num]; ok {
			return []int{mp[target-num], i}
		}
		mp[num] = i
	}
	return []int{}
}
