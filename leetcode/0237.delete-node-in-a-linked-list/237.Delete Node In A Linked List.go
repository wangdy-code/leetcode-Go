package leetcode

import "leetcode-go/leetcode/structures"

type ListNode = structures.ListNode

func deleteNode(node *ListNode) {
	t := node
	node = node.Next
	node.Next.Val = t.Val
}
