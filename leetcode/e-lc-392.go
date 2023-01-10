package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if s == t {
		return true
	}
	i := 0
	for j := range t {
		if t[j] == s[i] {
			if len(s) == i+1 {
				return true
			}
			i++
		}

	}
	return false
}
