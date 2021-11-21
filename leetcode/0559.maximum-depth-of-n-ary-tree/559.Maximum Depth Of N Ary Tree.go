package leetcode

import "leetcode-go/leetcode/structures"

type Node = structures.Node

/*深度优先*/
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	maxChildDepth := 0
	for _, child := range root.Children {
		if childDepth := maxDepth(child); childDepth > maxChildDepth {
			maxChildDepth = childDepth
		}
	}
	return maxChildDepth + 1
}

/*广度优先*/
func maxDepth1(root *Node) (ans int) {
	if root == nil {
		return
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		q := queue
		queue = nil
		for _, node := range q {
			queue = append(queue, node.Children...)
		}
		ans++
	}
	return
}
