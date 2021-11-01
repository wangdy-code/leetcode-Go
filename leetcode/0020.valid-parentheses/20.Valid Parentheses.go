package leetcode

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
