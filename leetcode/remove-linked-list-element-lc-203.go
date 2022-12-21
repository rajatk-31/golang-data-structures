package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	i, j := dummy, head
	for j != nil {
		if j.Val == val {
			i.Next = j.Next
			j = i.Next
		} else {
			i = i.Next
			j = j.Next
		}
	}
	return dummy.Next
}
