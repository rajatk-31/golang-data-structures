package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

func main() {
	tree := &Node{key: 100}
	tree.insert(50)
	tree.insert(120)
	fmt.Println(tree)
	fmt.Println(tree.search(213))
}

// insert
func (n *Node) insert(val int) {
	if n.key < val {
		//Move right
		if n.right == nil {
			n.right = &Node{key: val}
		} else {
			n.right.insert(val)
		}
	} else if n.key > val {
		//move left
		if n.left == nil {
			n.left = &Node{key: val}
		} else {
			n.left.insert(val)
		}
	}
}

// search
func (n *Node) search(val int) bool {
	if n == nil {
		return false
	}
	if n.key < val {
		//Move right
		return n.right.search(val)
	} else if n.key > val {
		//move left
		return n.left.search(val)
	}
	return true
	// if n.key == val {
	// 	return true
	// }
	// if n.key < val {
	// 	//Move right
	// 	if n.right != nil && n.right.key == val {
	// 		return true
	// 	} else if n.right == nil {
	// 		return false
	// 	} else {
	// 		return n.right.search(val)
	// 	}
	// } else if n.key > val {
	// 	//move left
	// 	if n.left != nil && n.left.key == val {
	// 		return true
	// 	} else if n.left == nil {
	// 		return false
	// 	} else {
	// 		return n.left.search(val)
	// 	}
	// }
	// return false
}
