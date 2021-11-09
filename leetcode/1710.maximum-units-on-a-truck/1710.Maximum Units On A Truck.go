package leetcode

type boxType struct {
	Type int
	Size int
}
type boxList []boxType

func maximumUnits(boxTypes [][]int, truckSize int) int {
	//按照boxTypes[i][1]排序
	for i := 0; i < len(boxTypes)-1; i++ {
		for j := 0; j < len(boxTypes)-1-i; j++ {
			if boxTypes[j][1] < boxTypes[j+1][1] {
				temp := boxTypes[j]
				boxTypes[j] = boxTypes[j+1]
				boxTypes[j+1] = temp
			}
		}
	}
	res := 0
	size := 0
	for _, row := range boxTypes {
		if size+row[0] <= truckSize {
			res += row[0] * row[1]
			size += row[0]
		} else {
			res += (truckSize - size) * row[1]
			size = truckSize
		}
	}
	return res
}
