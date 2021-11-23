package leetcode

func generateParenthesis(n int) (ans []string) {
	s := ""
	generate(s, n, &ans)
	return
}
func generate(s string, n int, res *[]string) {
	if len(s) == 2*n {
		if valid(s) {
			*res = append(*res, s)
		} else {
			s = s + "("
			generate(s, n, res)
			s = s[:len(s)-1]
			s += ")"
			generate(s, n, res)
			s = s[:len(s)-1]
		}
	}
}
func valid(s string) bool {
	bal := 0
	for _, v := range s {
		if v == '(' {
			bal++
		} else {
			bal--
		}
		if bal < 0 {
			return false
		}
	}
	return true
}
