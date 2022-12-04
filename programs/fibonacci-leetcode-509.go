package main

import "fmt"

func main() {
	fmt.Println(fib(6))
	fmt.Println(fibPrint(6))
}

func fibPrint(n int) []int {
	seq := []int{0, 1}
	for len(seq) <= n {
		last := seq[len(seq)-1]
		secondLast := seq[len(seq)-2]
		seq = append(seq, last+secondLast)
	}

	return seq
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	a, b := fib(n-1), fib(n-2)
	return a + b
}
