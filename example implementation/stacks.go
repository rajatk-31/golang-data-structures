package main

import "fmt"

type Stack struct {
	items []int
}

func main() {
	myStack := Stack{}
	myStack.Push(40)
	myStack.Push(50)
	myStack.Push(100)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack)
}

// Add value
func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

// Remove value
func (s *Stack) Pop() int {
	removed := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return removed
}
