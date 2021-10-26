package leetcode

/* hash表记录字符次数*/
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

/*数组记录次数*/
func canConstruct1(ransomNote string, magazine string) bool {
	pos := [26]int{}

	for _, v := range magazine {
		pos[v-'a']++
	}
	for _, v := range ransomNote {
		if pos[v-'a'] == 0 {
			return false
		} else {
			pos[v-'a']--
		}
	}
	return true
}
