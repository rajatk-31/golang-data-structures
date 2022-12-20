package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func preorderTraversal(root *TreeNode) []int {
	var final []int

	traverse(root, &final)
	return final
}

func traverse(node *TreeNode, output *[]int) {
	if node == nil {
		return
	}
	*output = append(*output, node.Val)
	traverse(node.Left, output)

	if node.Right != nil {
		traverse(node.Right, output)
	}

}
