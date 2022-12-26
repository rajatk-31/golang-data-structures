package main

func main() {

}

func canJump(nums []int) bool {
	lastGoodIndex := len(nums) - 1
	for i := lastGoodIndex; i >= 0; i-- {
		if i+nums[i] >= lastGoodIndex {
			lastGoodIndex = i
		}
	}
	return lastGoodIndex == 0
}
