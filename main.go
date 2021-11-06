package main

func main() {
	x:=[]int{3,0,1}
	missingNumber1(x)
}
func missingNumber1(nums []int) int {
	length := len(nums)
	indexNums := make([]int, length)
	for k, _ := range indexNums {
		indexNums[k] = k
	}
	for _, v := range nums {
		if v < length {
			indexNums = append(indexNums[:v], indexNums[v+1:]...)
		}
	}

	return indexNums[0]
}