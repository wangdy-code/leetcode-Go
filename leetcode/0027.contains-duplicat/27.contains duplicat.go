package leetcode

import "sort"

/*先对数组sort 然后遍历对前一位做对比*/
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

/* 使用哈希表 遍历数组 数组的value 充当 字典的key 给每个key都赋一个空结构体 当存在map[key] !=nil 条件为true*/
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
