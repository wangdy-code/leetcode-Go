package leetcode

import "leetcode-go/leetcode/structures"

type ListNode = structures.ListNode

/*遍历每一位数字 相加 大于0 下一位进1 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{Val: 0}
	current := head
	flag := false
	head1 := l1
	head2 := l2
	n1, n2 := 0, 0
	for head1 != nil || head2 != nil || carry != 0 {
		if head1 == nil {
			n1 = 0
		} else {
			n1 = head1.Val

			head1 = head1.Next
		}
		if head2 == nil {
			n2 = 0
		} else {
			n2 = head2.Val
			head2 = head2.Next
		}
		if flag {
			num := n1 + n2 + 1
			if num >= 10 {
				flag = true
				current.Next = &ListNode{Val: num - 10}
			} else {
				flag = false
				current.Next = &ListNode{Val: num}
			}
		} else {
			num := n1 + n2
			if num >= 10 {
				flag = true
				current.Next = &ListNode{Val: num - 10}
			} else {
				flag = false
				current.Next = &ListNode{Val: num}
			}
		}
		current = current.Next
		carry = (n1 + n2 + carry) / 10

	}
	return head.Next
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}
