package leetcode

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	}
	return 2 + min(integerReplacement(n/2), integerReplacement(n/2+1))
}

var memo map[int]int

func _integerReplacement(n int) (res int) {
	if n == 1 {
		return 0
	}
	if res, ok := memo[n]; ok {
		return res
	}
	defer func() { memo[n] = res }()
	if n%2 == 0 {
		return 1 + _integerReplacement(n/2)
	}
	return 2 + min(_integerReplacement(n/2), _integerReplacement(n/2+1))
}

func integerReplacement2(n int) (res int) {
	memo = map[int]int{}
	return _integerReplacement(n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func integerReplacement1(n int) (ans int) {
	for n != 1 {
		switch {
		case n%2 == 0:
			ans++
			n /= 2
		case n%4 == 1:
			ans += 2
			n /= 2
		case n == 3:
			ans += 2
			n = 1
		default:
			ans += 2
			n = n/2 + 1
		}
	}
	return
}
