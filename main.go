package main

import (
	"leetcode-go/leetcode/structures"
	"sort"
)

type myStruct struct {
}
type ListNode = structures.ListNode

func main() {
	boxTypes := [][]int{
		{5, 10}, {2, 5}, {4, 7}, {3, 9},
	}
	maximumUnits(boxTypes, 10)
}
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	cMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if cMap[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != cMap[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
func canConstruct(ransomNote string, magazine string) bool {
	strMap := make(map[string]int)
	for _, v := range magazine {
		if _, has := strMap[string(v)]; has {
			continue
		} else {
			strMap[string(v)] = 1
		}
	}
	for _, v := range ransomNote {
		if _, has := strMap[string(v)]; !has {
			return false
		} else {
			delete(strMap, string(v))
		}
	}
	return true
}
func firstUniqChar(s string) int {
	charTimes := make(map[string]int)
	for _, v := range s {
		charTimes[string(v)]++
	}
	for k, v := range s {
		if charTimes[string(v)] <= 1 {
			return k
		}
	}
	return -1
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
func maximumUnits(boxTypes [][]int, truckSize int) int {
	//按照boxTypes[i][1]排序
	for i := 0; i < len(boxTypes)-1; i++ {
		for j := 0; j < len(boxTypes)-1-i; j++ {
			if boxTypes[j][1] < boxTypes[j+1][1] {
				temp := boxTypes[j]
				boxTypes[j] = boxTypes[j+1]
				boxTypes[j+1] = temp
			}
		}
	}
	res := 0
	size := 0
	for _, row := range boxTypes {
		if size+row[0] <= truckSize {
			res += row[0] * row[1]
			size += row[0]
		} else {
			res += (truckSize - size) * row[1]
			size = truckSize
		}
	}
	return res
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
