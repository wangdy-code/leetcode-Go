package leetcode

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]int)
	maxLen := 0
	end := -1
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(charMap, s[i-1])
		}
		for end+1 < len(s) && charMap[s[end+1]] == 0 {
			charMap[s[end+1]]++
			end++
		}
		maxLen = max(maxLen, end-i+1)
	}
	return maxLen
}
func lengthOfLongestSubstring_(s string) int {
	charMap := make(map[byte]int)
	maxLen := 0
	n := len(s)
	start, end := 0, 0
	for end < n {
		if start != 0 {
			delete(charMap, s[start-1])
		}
		if charMap[s[end]] != 0 {

			maxLen = max(maxLen, end-start)
		} else {
			charMap[s[end]]++
		}
		start = end + 1

	}
	return maxLen
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
