package leetcode

import "math"

const mod = 1337

func superPow(a int, b []int) int {
	ans := 1
	for _, e := range b {
		ans = int(math.Pow(float64(ans), 10)*math.Pow(float64(a), float64(e))) % mod
	}
	return ans
}
