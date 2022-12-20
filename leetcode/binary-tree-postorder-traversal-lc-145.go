package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func postorderTraversal(root *TreeNode) []int {
	var final []int

	traverse(root, &final)
	return final
}

func traverse(node *TreeNode, output *[]int) {
	if node == nil {
		return
	}

	traverse(node.Left, output)

	traverse(node.Right, output)

	*output = append(*output, node.Val)

}
