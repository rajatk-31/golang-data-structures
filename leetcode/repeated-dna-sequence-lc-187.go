package main

func main() {

}

func findRepeatedDnaSequences(s string) []string {

	all := map[string]int{}
	final := []string{}
	for i := 0; i <= len(s)-10; i++ {
		seq := s[i : i+10]
		if _, ok := all[seq]; ok {
			all[seq] += 1
		} else {
			all[seq] = 1
		}
	}
	for k, v := range all {
		if v > 1 {
			final = append(final, k)
		}
	}
	return final
}
