package leetcode

func maxPower(s string) (ans int) {
	if len(s) == 1 {
		return 1
	}
	left, right, length := 0, 1, 1
	for left < len(s) && right < len(s) {
		if s[left] == s[right] {
			length++
		} 
        if length > ans {
				ans = length
				
			}
            if s[left]!=s[right]{
                length = 1
            }
		left++
		right++
	}
	return 
}
func maxPower_(s string) int {
    ans, cnt := 1, 1
    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1] {
            cnt++
            if cnt > ans {
                ans = cnt
            }
        } else {
            cnt = 1
        }
    }
    return ans
}