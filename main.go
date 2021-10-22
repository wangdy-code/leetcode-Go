package main

import (
	"leetcode-go/leetcode/structures"
	"sort"
)

type ListNode = structures.ListNode

func main() {
	num1 := []int{1, 2}
	nums2 := []int{3, 4}

	println(findMedianSortedArrays1(num1, nums2))
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
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sort.Ints(nums1)
	sort.Ints(nums2)
	m, n := len(nums1), len(nums2)
	index1, index2 := 0, 0
	res := make([]int, 0, m+n)
	for {
		if index1 == m {
			res = append(res, nums2[index2:]...)
			break
		}
		if index2 == n {
			res = append(res, nums1[index1:]...)
			break
		}
		if nums1[index1] < nums2[index2] {
			res = append(res, nums1[index1])
			index1++
		} else {
			res = append(res, nums2[index2])
			index2++
		}
	}
	if (m+n)%2 != 0 {
		return float64(res[(m+n)/2])
	} else {
		return float64((res[(m+n)/2] + res[(m+n)/2-1])) / float64(2)
	}
}
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {

	m, n := len(nums1), len(nums2)
	index, index1, index2 := 0, 0, 0
	isOdd := false
	if (m+n)%2 == 0 {
		isOdd = true
	}
	for {
		if index1 == m {
			if index == (m+n)/2 {
				if isOdd {
					if index2 == 0 {
						return (float64(nums2[index2]) + float64(nums1[m-1])) / 2
					} else {
						return (float64(nums2[index2]) + float64(nums2[index2-1])) / 2
					}
				} else {
					return float64(nums2[index2])
				}
			}
			index2++
			index++
			continue
		}
		if index2 == n {
			if index == (m+n)/2 {
				if isOdd {
					if index2 == 0 {
						return (float64(nums1[index1]) + float64(nums2[n-1])) / 2
					} else {
						return (float64(nums1[index1]) + float64(nums1[index1-1])) / 2
					}
				} else {
					return float64(nums1[index1])
				}
			}
			index1++
			index++
			continue
		}
		if nums1[index1] < nums2[index2] {
			index1++
			index++
		} else {
			index2++
			index++
		}
	}

}
