package main

func main() {

}

func runningSum(nums []int) []int {
	final := []int{}
	for i := range nums {
		// fmt.Println(v)
		if i == 0 {
			final = append(final, nums[i])

		} else {
			final = append(final, nums[i]+final[i-1])
		}
	}
	return final
}
