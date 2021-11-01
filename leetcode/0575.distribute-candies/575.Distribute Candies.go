package leetcode

import "sort"

func distributeCandies(candyType []int) int {
	sort.Ints(candyType)
	averageNums := len(candyType) / 2
	tMap := make([]int, len(candyType))
	typeNums := 0
	for _, v := range candyType {
		if tMap[v] > 0 {
			tMap[v]++
		} else {
			tMap[v] = 1
			typeNums++
		}
	}
	if typeNums >= averageNums {

		return averageNums
	} else {
		return typeNums
	}

}
