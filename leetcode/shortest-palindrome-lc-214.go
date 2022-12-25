package main

func main() {

}

func shortestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	index := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[index] == s[i] {
			index++
		}
	}
	if index == len(s) {
		return s
	}
	return reverse(s[index:]) + shortestPalindrome(s[:index]) + s[index:]
}

func reverse(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
