package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		/*判断目标值 如果不是  v作为键 索引作为字典的值
		  第二次遍历是目标值减去遍历值 字典中是否有这个键
		  是则返回这连个值*/
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}

/*简单解法直接遍历两次 时间复杂度n^2*/
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
