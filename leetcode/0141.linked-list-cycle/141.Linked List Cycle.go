package leetcode

import "leetcode-go/leetcode/structures"

type ListNode = structures.ListNode

/*使用hash表记录节点 判断节点是否存在 存在即为环形链表*/
func hasCycle(head *ListNode) bool {
	nodeMap := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := nodeMap[head]; ok {
			return true
		}
		nodeMap[head] = struct{}{}
		head = head.Next
	}
	return false
}

/*使用快慢指针 当快指针追上慢指针时 证明链表存在环形*/
func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if slow == nil || fast == nil {
			return false
		}
		slow = slow.Next
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
	}
	return true
}
