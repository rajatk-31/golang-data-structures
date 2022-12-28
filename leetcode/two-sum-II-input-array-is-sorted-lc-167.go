package main

func main() {

}

func twoSum(numbers []int, target int) []int {
	test := map[int]int{}
	for i, v := range numbers {
		diff := target - v

		if _, ok := test[diff]; ok {
			return []int{test[diff] + 1, i + 1}
		} else {
			test[v] = i
		}
	}
	return []int{}
}
