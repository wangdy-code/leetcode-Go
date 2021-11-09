package leetcode

/*二分查找*/
func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := (left + right) / 2
		if mid^2 < num {
			left = mid + 1
		} else if mid^2 > num {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}

/*牛顿迭代*/
func isPerfectSquare1(num int) bool {
	x0 := float64(num)
	for {
		x1 := (x0 + float64(num)/x0) / 2
		if x0-x1 < 1e-6 {
			x := int(x0)
			return x*x == num
		}
		x0 = x1
	}
}
