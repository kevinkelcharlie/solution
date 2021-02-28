/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.com/problems/invert-binary-tree/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// Invert all right node
	right := invertTree(rozt.Right)

	// Invert all left node
	left := invertTree(root.Left)

	// Swap Node
	root.Left = right
	root.Right = left
	return root
}

