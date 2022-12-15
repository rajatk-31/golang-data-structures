package main

import (
	"math"
)

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tempList := &ListNode{}
	current := tempList
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = int(math.Floor(float64(sum) / 10))
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}
	return tempList.Next
}
