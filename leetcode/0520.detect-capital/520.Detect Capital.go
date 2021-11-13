package leetcode

func detectCapitalUse(word string) bool {
	if len(word) < 2 {
		return true
	}
	if 'a' <= word[0] && word[0] <= 'z' {
		for i := 1; i < len(word); i++ {
			if word[i] <= 'Z' && word[i] >= 'A' {
				return false
			}
		}
	} else {
		if word[1] <= 'Z' && word[1] >= 'A' {
			for i := 2; i < len(word); i++ {
				if 'a' <= word[i] && word[i] <= 'z' {
					return false
				}
			}
		} else {
			for i := 2; i < len(word); i++ {
				if word[i] <= 'Z' && word[i] >= 'A' {
					return false
				}
			}
		}

	}
	return true
}
