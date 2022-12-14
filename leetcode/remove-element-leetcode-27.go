package main

import "fmt"

func main() {
	arr := []int{3, 2, 2, 3}
	fmt.Println(removeElement(arr, 3))
}

func removeElement(nums []int, val int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
