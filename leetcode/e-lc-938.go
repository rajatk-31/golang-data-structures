package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	if root == nil {
		return sum
	}
	c(root, low, high, &sum)
	// if root.Val>=low && root.Val<=high{
	// sum+= root.Val
	// }
	// if root.Right!=nil{
	//     rangeSumBST(root.Right, low , high)
	// }
	// if root.Left!=nil{
	//     rangeSumBST(root.Left, low , high)
	// }
	return sum

}
func c(root *TreeNode, low int, high int, sum *int) {
	if root.Val >= low && root.Val <= high {
		*sum += root.Val
	}
	if root.Right != nil {
		c(root.Right, low, high, sum)
	}
	if root.Left != nil {
		c(root.Left, low, high, sum)
	}
}
