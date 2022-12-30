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
func middleNode(head *ListNode) *ListNode {
	//    l := 0
	//     tmp := head
	//     for tmp != nil {
	//         tmp = tmp.Next
	//         l++
	//     }
	//     for i := 0; i < l / 2; i++ {
	//         head = head.Next
	//     }
	//     return head
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow

}
