package main

import (
	"leetcode-go/leetcode/structures"
)

type ListNode = structures.ListNode

func main() {
	nums := []int{-1, 2, -2, 4}
	println(maxSubArray(nums))
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
