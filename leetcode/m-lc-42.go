package main

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
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return nil
	}
	fast := head
	slow := head
	pointer := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast != slow {
		return nil
	}
	for pointer != slow {
		pointer = pointer.Next
		slow = slow.Next

	}
	return slow
}
