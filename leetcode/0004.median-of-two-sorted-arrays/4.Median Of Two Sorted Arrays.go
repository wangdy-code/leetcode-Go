package leetcode

import "sort"

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
