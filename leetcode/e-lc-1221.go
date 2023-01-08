package main

func main() {

}

func balancedStringSplit(s string) int {
	count := 0
	ans := 0

	for _, r := range s {
		if r == 'R' {
			count++
		} else {
			count--
		}
		if count == 0 {
			ans++
		}
	}

	return ans
}
