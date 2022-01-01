package main

func truncateSentence(s string, k int) (ans string) {
	str := ""
	stringList := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			str += string(s[i])
		}
		if s[i] == ' ' || i == len(s)-1 {
			stringList = append(stringList, str)
			str = ""
		}
	}
	for i, strl := range stringList {
		if i == k {
			break
		}
		ans += strl
		if i != k-1 {
			ans += " "
		}

	}
	return
}
