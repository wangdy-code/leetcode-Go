package leetcode

/*暴力解法*/
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	maxLen := 1
	begin := 0
	var dp [][]bool

	row := []bool{}
	for i := 0; i < n; i++ {
		row = append(row, false)
	}
	for i := 0; i < n; i++ {
		dp = append(dp, row)
		dp[i][i] = true
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {

			if j-i+1 > maxLen && validPalindromic(s, i, j) {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]

}

/*中心扩展*/
func longestPalindrome1(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start = left1
			end = right1
		}
		if right2-left2 > end-start {
			start = left2
			end = right2
		}
	}
	return s[start : end+1]

}
func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
func validPalindromic(str string, left, right int) bool {
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
