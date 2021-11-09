package leetcode

import "fmt"

func getHint(secret string, guess string) string {
	cMap := make(map[byte]int)
	A, B := 0, 0
	for k := 0; k < len(secret); k++ {
		if secret[k] == guess[k] {
			A++
			guess = guess[:k] + guess[k+1:]
			secret = secret[:k] + guess[k+1:]
			k--
			continue
		}
		if _, has := cMap[secret[k]]; has {
			cMap[secret[k]]++
		} else {
			cMap[secret[k]] = 1
		}
	}

	for k, _ := range guess {
		if _, has := cMap[guess[k]]; has && cMap[guess[k]] > 0 {
			B++
			cMap[guess[k]]--
		}
	}
	return fmt.Sprintf("%dA%dB", A, B)
}
func getHint1(secret, guess string) string {
	bulls := 0
	var cntS, cntG [10]int
	for i := range secret {
		if secret[i] == guess[i] {
			bulls++
		} else {
			cntS[secret[i]-'0']++
			cntG[guess[i]-'0']++
		}
	}
	cows := 0
	for i := 0; i < 10; i++ {
		cows += min(cntS[i], cntG[i])
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
