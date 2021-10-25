package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	strMap := make(map[string]int)
	for _, v := range magazine {
		if _, has := strMap[string(v)]; has {
			strMap[string(v)]++
		} else {
			strMap[string(v)] = 1
		}
	}
	for _, v := range ransomNote {
		if _, has := strMap[string(v)]; !has {
			return false
		} else {
			strMap[string(v)]--
			if strMap[string(v)] == 0 {

				delete(strMap, string(v))
			}
		}
	}
	return true
}
