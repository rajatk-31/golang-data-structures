package main

func main() {

}

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}
	m := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			lastIndex := len(stack) - 1
			if lastIndex == -1 {
				return false
			}
			lastOpenParenthese := stack[lastIndex]
			if v == m[lastOpenParenthese] {
				stack = stack[:lastIndex]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
