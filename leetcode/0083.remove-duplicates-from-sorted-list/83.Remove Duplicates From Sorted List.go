package leetcode

import "leetcode-go/leetcode/structures"

type ListNode = structures.ListNode

/*判断当前节点和下一个节点是否相等*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
