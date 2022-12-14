package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 16}
	node4 := &node{data: 11}
	node5 := &node{data: 7}
	node6 := &node{data: 2}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)

	myList.print()
	myList.delete(100)
	myList.delete(2)
	myList.delete(48)
	myList.print()

}

/*
	Append at start
*/
// if use without * then we work on a copy of that else original will be changeed
func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

/*
Print Data
*/
func (l linkedList) print() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d\t", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println("All done")
}

/*
Delete Data
*/
func (l *linkedList) delete(val int) {

	if l.head.data == val {
		l.head = l.head.next
		l.length--
		return
	}
	if l.length == 0 {
		fmt.Println("Nothing to delete")
		return
	}
	toDelete := l.head
	for toDelete.next.data != val {
		if toDelete.next.next == nil {
			return
		}
		toDelete = toDelete.next
	}
	toDelete = toDelete.next.next
	l.length--

}
