package main

func main() {

}

func rotate(nums []int, k int) {
	i := int32(len(nums) - k%len(nums))
	copy(nums, append(nums[i:], nums[:i]...))
}
