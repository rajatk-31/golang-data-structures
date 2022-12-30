package main

import "fmt"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := 0
	tmp := head
	for tmp != nil {
		tmp = tmp.Next
		l++
	}

	if l == n {
		return head.Next
	}
	toRemove := l - n
	it := head
	for toRemove >= 0 {
		fmt.Println(it.Val, toRemove)
		if toRemove == 1 {
			it.Next = it.Next.Next
			break
		}
		it = it.Next
		toRemove--
	}
	return head
}
