package offer

import "math"

func divide(a int, b int) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	sign := 1
	if (a > 0 && b < 0) || a < 0 && b > 0 {
		sign = -1
	}
	a, b = int(math.Abs(float64(a))), int(math.Abs(float64(b)))
	res := 0
	for a >= b {
		a -= b
		res++
	}

	return sign * res
}
