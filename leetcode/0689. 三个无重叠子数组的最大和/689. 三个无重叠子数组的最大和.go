package leetcode

func maxSumOfThreeSubarrays(nums []int, k int) (ans []int) {
	var sum1, maxSum1, maxSum1Index int
	var sum2, maxSum12, maxSum12Index int
	var sum3, maxTotal int
	for i := k * 2; i < len(nums); i++ {
		sum1 += nums[i-k*2]
		sum2 += nums[i-k]
		sum3 += nums[i]
		if i > k*3-1 {
			if sum1 > maxSum1 {
				maxSum1 = sum1
				maxSum1Index = i - k*3 + 1
			}
			if maxSum1+sum2 > maxSum12 {
				maxSum12 = maxSum1 + sum2
				maxSum12Index = i - k*2 + 1
			}
			if maxSum12+sum3 > maxTotal {
				maxTotal = maxSum12 + sum3
				ans = []int{maxSum1Index, maxSum12Index, i - k + 1}
			}
			sum1 -= nums[i-k*3+1]
			sum2 -= nums[i-k*2+1]
			sum3 -= nums[i-k+1]
		}
	}
	return
}
