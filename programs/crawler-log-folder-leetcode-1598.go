package main

import "fmt"

func main() {
	fmt.Println(minOperations([]string{"d1/", "d2/", "../", "d21/", "./"}))
}

func minOperations(logs []string) int {
	var no int

	for i := 0; i < len(logs); i++ {
		if logs[i] == "../" {
			if no > 0 {
				no--
			}
		} else if logs[i] == "./" {
			//do nothing
		} else {
			no++
		}
	}

	return no
}
