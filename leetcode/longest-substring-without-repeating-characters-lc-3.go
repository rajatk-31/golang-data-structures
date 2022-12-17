package main

import "math"

func main() {

}

func lengthOfLongestSubstring(s string) int {
	result := -1
	dict := make(map[byte]int)
	d := -1

	for i := 0; i < len(s); i++ {
		if v, ok := dict[s[i]]; !ok {
			dict[s[i]] = i
			d++
		} else {
			d = int(math.Min(float64(i-v-1), float64(d+1)))
			dict[s[i]] = i
		}
		result = int(math.Max(float64(result), float64(d)))
	}
	return result + 1

}
