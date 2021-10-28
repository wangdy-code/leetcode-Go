package leetcode

import "leetcode-go/leetcode/structures"

type ListNode = structures.ListNode

/*迭代*/
func reverseList(head *ListNode) *ListNode {
	var cur *ListNode
	cur = nil
	pre := head
	for pre != nil {
		t := pre.Next
		pre.Next = cur
		cur = pre
		pre = t
	}
	return cur
}

/*递归*/
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}
