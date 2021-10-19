package leetcode

import "sort"

/*将数组便利 存入hash表 遍历数组2查看次数*/
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	m := map[int]int{}
	for _, v := range nums1 {
		m[v]++
	}
	intsection := []int{}
	for _, v := range nums2 {
		if m[v] > 0 {
			intsection = append(intsection, v)
			m[v]--
		}
	}
	return intsection
}
func intersect1(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	length1, length2 := len(nums1), len(nums2)
	index1, index2 := 0, 0
	intsection := []int{}
	for index1 < length1 && index2 < length2 {
		if nums1[index1] < nums2[index2] {
			index1++
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else {
			intsection = append(intsection, nums1[index1])
			index1++
			index2++
		}
	}
	return intsection
}
