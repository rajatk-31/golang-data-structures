package main

type Node struct {
	Val      int
	Children []*Node
}

func main() {

}

func postorder(root *Node) []int {
	var final []int

	traverse(root, &final)
	return final
}

func traverse(node *Node, output *[]int) {
	if node == nil {
		return
	}
	for i := 0; i < len(node.Children); i++ {
		traverse(node.Children[i], output)
	}

	*output = append(*output, node.Val)

}
