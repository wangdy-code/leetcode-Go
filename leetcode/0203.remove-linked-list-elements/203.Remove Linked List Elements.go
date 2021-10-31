package leetcode

import "leetcode-go/leetcode/structures"

type ListNode = structures.ListNode

/*递归*/
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

/*迭代*/
func removeElements1(head *ListNode, val int) *ListNode {
	prev := &ListNode{Next: head}
	for tmp := prev; tmp.Next != nil; {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return prev.Next
}
