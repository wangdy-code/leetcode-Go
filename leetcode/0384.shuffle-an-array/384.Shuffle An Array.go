package leetcode

import (
	"math/rand"
)

type Solution struct {
	nums, original []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums, original: append([]int{}, nums...)}
}

func (this *Solution) Reset() []int {
	copy(this.nums, this.original)
	return this.nums

}

func (this *Solution) Shuffle() []int {
	n := len(this.nums)
	for k, _ := range this.nums {
		x := rand.Intn(n)
		t := this.nums[k]
		this.nums[k] = this.nums[x]
		this.nums[x] = t
	}
	return this.nums
}
