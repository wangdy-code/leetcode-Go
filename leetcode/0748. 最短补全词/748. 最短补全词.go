package leetcode

import (
	"unicode"
)

func shortestCompletingWord(licensePlate string, words []string) (ans string) {
	nums := [26]int{0}
	for _, v := range licensePlate {
		if unicode.IsLetter(v) {
			nums[unicode.ToLower(v)-'a']++
		}
	}
	for _, v := range words {
		temp := nums
		for _, ch := range v {
			if temp[ch-'a'] > 0 {
				temp[ch-'a']--
			} else {
				break
			}
			if ans == "" || len(v) < len(ans) {
				ans = v
			}
		}
	}
	return
}
