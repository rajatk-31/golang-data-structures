package main

func main() {

}

func sortedSquares(nums []int) []int {
	var ptr1, ptr2 int = 0, len(nums) - 1
	var currentIndex int = len(nums) - 1
	result := make([]int, len(nums))
	for ptr2 >= ptr1 {
		val1, val2 := nums[ptr1], nums[ptr2]
		sqr1, sqr2 := val1*val1, val2*val2
		if sqr1 >= sqr2 {
			result[currentIndex] = sqr1
			ptr1++
		} else /* if sqr1 < sqr2 */ {
			result[currentIndex] = sqr2
			ptr2--
		}
		currentIndex--
	}
	return result
}
