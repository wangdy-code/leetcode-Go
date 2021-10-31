package leetcode

/*不需要额外的空间 时间和数组长度有关*/
func singleNumber(nums []int) int {
	singlenumber := 0
	for _, v := range nums {
		singlenumber ^= v //对数组中每一个元素进行异或运算  任意元素对0做异或运算等于它本身 对它本身做该运算 为0
	}
	return singlenumber
}
