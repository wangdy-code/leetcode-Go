package leetcode

func lastRemaining(n int) int {
	first := 1
	k, cnt, step := 0, n, 1
	for cnt > 1 {
		if k%2 == 0 {
			first += step
		} else {
			if cnt%2 == 1 {
				first += step
			}
		}
		k++
		cnt >>= 1
		step <<= 1
	}
	return first
}
