package leetcode

import "leetcode-go/leetcode/structures"

type TreeNode = structures.TreeNode

func findTilt(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sumLeft := dfs(node.Left)
		sumRight := dfs(node.Right)
		ans += abs(sumLeft - sumRight)
		return sumLeft + sumRight + node.Val
	}
	dfs(root)
	return
}
func abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}
