package leetcode

import (
	"math"
)

func reverse(x int) (res int) {

	for x != 0 {
		if res > math.MaxInt32/10 || res < math.MinInt32/10 {
			return 0
		}
		dig := x % 10
		x /= 10
		res = res*10 + dig
	}
	return
}
