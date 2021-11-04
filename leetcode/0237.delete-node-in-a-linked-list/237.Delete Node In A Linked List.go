package leetcode

import "leetcode-go/leetcode/structures"

type ListNode = structures.ListNode

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
