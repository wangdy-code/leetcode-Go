package leetcode

func missingNumber(nums []int) int {
	length := len(nums)
	indexNums := make([]int, length)
	for _, v := range nums {
		if v < length {
			indexNums[v] = 1
		}
	}
	for k, v := range indexNums {
		if v == 0 {
			return k
		}
	}
	return length
}
func missingNumber1(nums []int) int {
	length := len(nums)
	sum := length * (length + 1) / 2
	arrSum := 0
	for _, v := range nums {
		arrSum += v
	}
	return sum - arrSum
}
