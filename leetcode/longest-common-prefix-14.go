package main

import (
	"fmt"
)

func main() {
	fmt.Println("ans", longestCommonPrefix([]string{"c", "acc", "ccc"}))
}

func longestCommonPrefix(strs []string) string {
	// st := strs[0]

	// for i := 1; i < len(strs); i++ {
	// 	found := false
	// 	for !found && len(st) > 0 {

	// 		if len(st) > len(strs[i]) {
	// 			st = st[:len(strs[i])]
	// 		}

	// 		toCompare := strs[i][:len(st)]
	// 		if !(toCompare == st) {
	// 			st = st[:len(st)-1]
	// 		} else {
	// 			found = true
	// 		}
	// 	}
	// }
	// return st
	index := 0
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	for {
		for i := 1; i < len(strs); i++ {
			if index >= len(strs[i-1]) || index >= len(strs[i]) || strs[i-1][index] != strs[i][index] {
				return strs[0][:index]
			}
		}
		index++
	}
	// return strs[0][:index]

}
