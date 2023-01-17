package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	finalArr := []int{}
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	start := -1
	end := -1
	for _, v := range arr2 {
		for i, v2 := range arr1 {
			if v2 == v {
				if start == -1 {
					start = i
					end = i
				} else {
					end = i
				}

				finalArr = append(finalArr, v2)
				// arr1=append(arr1[:i], arr1[i+1:]...)
			}
		}
		if start != -1 {
			arr1 = append(arr1[:start], arr1[end+1:]...)
			start = -1
			end = -1
		}
	}
	finalArr = append(finalArr, arr1...)
	return finalArr
}
