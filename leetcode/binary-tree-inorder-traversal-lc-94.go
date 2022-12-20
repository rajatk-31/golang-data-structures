package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var final []int

func inorderTraversal(root *TreeNode) []int {
	//left, root,right

	var final []int

	traverse(root, &final)
	return final

}

func traverse(node *TreeNode, output *[]int) {
	if node == nil {
		return
	}
	traverse(node.Left, output)
	*output = append(*output, node.Val)
	if node.Right != nil {
		traverse(node.Right, output)
	}

}
