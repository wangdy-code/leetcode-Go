package leetcode

func maxArea(height []int) int {
	ans, left := 0, 0
	right := len(height) - 1
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		ans = max(ans, area)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
