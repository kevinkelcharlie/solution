/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, isBalance := getHeightAndIsBalance(root)
	return isBalance

}

func getHeightAndIsBalance(root *TreeNode) (int, bool) {
	// If leaf is nil
	if root == nil {
		return 0, true
	}

	// Check left node
	left, lBalance := getHeightAndIsBalance(root.Left)
	if !lBalance {
		return 0, false
	}
	// Check Right node
	right, rBalance := getHeightAndIsBalance(root.Right)
	if !rBalance {
		return 0, false
	}

	isBalance := getDiff(left, right) < 2
	return getMax(left, right) + 1, isBalance

}

func getMax(left, right int) int {
	max := left
	if right > left {
		max = right
	}
	return max
}

func getDiff(left, right int) int {
	return int(math.Abs(float64(left) - float64(right)))
}
