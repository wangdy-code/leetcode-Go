package leetcode

import "strconv"

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
func isPalindrome1(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}
	return revertedNumber == x || x == revertedNumber/10
}
func isPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	str := strconv.FormatInt(int64(x), 10)
	for i := 0; i < len(str)/2+1; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
