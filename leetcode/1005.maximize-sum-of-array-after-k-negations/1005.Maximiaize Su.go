package leetcode

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) (ans int) {
	sort.Ints(nums)
	cnt := 0
	zero := false
	idx := 0
	fNums := []int{}
	for k, v := range nums {
		if v < 0 {
			fNums = append(fNums, v)
			cnt++
		}
		if v == 0 {
			zero = true
		}
		if math.Abs(float64(nums[k])) < math.Abs(float64(nums[idx])) {
			idx = k
		}
	}
	if k <= cnt {
		for i := 0; i < k; i++ {
			nums[i] = -nums[i]
		}
		for _, v := range nums {
			ans += v
		}
	} else {
		if (k-cnt)%2 == 0 || zero {
			for i := 0; i < cnt; i++ {
				fNums[i] = -fNums[i]
			}
			for _, v := range fNums {
				ans += v
			}
			for _, v := range nums[len(fNums):] {
				ans += v
			}
		} else if (k-cnt)%2 != 0 && !zero {
			for i := 0; i < cnt; i++ {
				nums[i] = -nums[i]
			}
			nums[idx] = -nums[idx]
			for _, v := range nums {
				ans += v
			}
		}

	}

	return
}
