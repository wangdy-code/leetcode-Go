package main

import (
	"leetcode-go/leetcode/structures"
)

type ListNode = structures.ListNode

func main() {
	set := map[int]struct{}{}
	println(set[0])
	set[1] = struct{}{}
	println(set[1])

}
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
func containsDuplicate1(nums []int) bool {
	set := map[int]struct{}{}
	for _, v := range nums {
		if _, has := set[v]; has {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}
