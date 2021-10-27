package main

import "leetcode-go/leetcode/structures"

type ListNode = structures.ListNode

/*使用递归*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}

}

/*使用迭代*/
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := new(ListNode)
	prev := prehead
	for {
		if l1 == nil {
			prev.Next = l2
			break
		}
		if l2 == nil {
			prev.Next = l1
			break
		}
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	return prehead.Next
}
