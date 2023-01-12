package main

func pivotInteger(n int) int {
	sum := 0
	reverse_sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	for i := 1; i <= n; i++ {
		if sum == reverse_sum+i {
			return i
		}
		sum -= i
		reverse_sum += i
	}
	return -1
}
