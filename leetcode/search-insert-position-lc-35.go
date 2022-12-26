package main

func main() {

}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	for left != right {
		mid := ((left + right) / 2)

		if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}

	}
	if right != 0 && right <= len(nums) && nums[right-1] == target {
		right--
	}
	return right
}
