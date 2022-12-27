package main

func main() {

}

func moveZeroes(nums []int) {
	i := 0
	zero := 0
	for i < len(nums) {
		if len(nums)-zero == i+1 {
			break
		}
		if nums[i] == 0 {
			x := append(nums[:i], nums[i+1:]...)
			zero++
			x = append(x, 0)
			copy(nums, x)
		} else {
			i++
		}
	}
}
