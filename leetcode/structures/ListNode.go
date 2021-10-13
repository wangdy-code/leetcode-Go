package structures

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func List2Ints(head *ListNode) []int {
	limit := 100
	times := 0
	res := []int{}
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	l := &ListNode{}
	t := l
	for _, value := range nums {
		t.Next = &ListNode{Val: value}
		t = t.Next
	}
	return l.Next
}

// GetNodeWith returns the first node with val
func (l *ListNode) GetNodeWith(val int) *ListNode {
	res := l
	for res != nil {
		if res.Val == val {
			break
		}
		res = res.Next
	}
	return res
}

// Ints2ListWithCycle returns a list whose tail point to pos-indexed node
// head's index is 0
// if pos = -1, no cycle
func Ints2ListWithCycle(nums []int, pos int) *ListNode {
	head := Ints2List(nums)
	if pos == -1 {
		return head
	}
	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}
	tail := c
	if tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c
	return tail
}
