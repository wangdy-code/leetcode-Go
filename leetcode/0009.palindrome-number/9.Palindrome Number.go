package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	nums := []int{}
	for x != 0 {
		nums = append(nums, x%10)
		x = x / 10
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[len(nums)-1-i] {
			return false
		}
	}
	return true
}
