package leetcode

func isAnagram(s string, t string) bool {
	pos1 := [26]byte{}
	pos2 := [26]byte{}
	for _, c := range s {
		pos1[c-'a']++
	}
	for _, c := range t {
		pos2[c-'a']++
	}
	for i := 0; i < len(pos1); i++ {
		if pos1[i] != pos2[i] {
			return false
		}
	}
	return true
}
