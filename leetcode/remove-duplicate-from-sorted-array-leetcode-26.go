package main

import "fmt"

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(arr))
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			nums[i+1] = nums[j]
			i++
			j++
		}
	}
	return i + 1

}
