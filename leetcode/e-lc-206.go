package main

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
func reverseList(head *ListNode) *ListNode {
	var temp *ListNode
	if head == nil {
		return head
	}
	for head != nil {
		des := &ListNode{Val: head.Val}
		des.Next = temp
		temp = des
		head = head.Next
	}
	return temp
}
