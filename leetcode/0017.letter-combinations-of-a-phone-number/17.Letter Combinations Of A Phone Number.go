package leetcode

/*DFS回溯算法*/
var phoneMap map[byte][]byte = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) (ans []string) {
	n := len(digits)
	if n == 0 {
		return
	}
	sb := ""
	dfs(digits, 0, n, sb, ans)
}
func dfs(ds string, i, n int, sb string, ans []string) {
	if i == n {
		ans = append(ans, sb)
		return
	}
	all:=phoneMap[ds]
	for _, c := range ds[i : i+1] {
		sb = sb + string(c)
		dfs(ds, i+1, n, sb, ans)
		sb = sb[:len(sb)-2]
	}
}
