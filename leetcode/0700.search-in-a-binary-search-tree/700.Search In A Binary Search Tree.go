package leetcode

import "leetcode-go/leetcode/structures"

type TreeNode = structures.TreeNode

/*递归*/
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}

/*迭代*/
func searchBST_(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if val < root.Val {
			root = root.Left
		} else if val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
