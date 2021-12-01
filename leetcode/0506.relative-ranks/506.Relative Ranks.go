package leetcode

import (
	"sort"
	"strconv"
)

type pair struct {
	score, index int
}

var desc = [3]string{"Gold Medal", "Silver Medal", "Bronze Medal"}

func findRelativeRanks(score []int) (ans []string) {
	n := len(score)
	arr := make([]pair, n)
	for i, s := range score {
		arr[i] = pair{s, i}
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].score > arr[j].score })
	for i, _ := range arr {
		if i < 3 {
			ans = append(ans, desc[i])
		} else {
			ans = append(ans, strconv.Itoa(i+1))
		}
	}
	return
}
func findRelativeRanks_(score []int) []string {
	n := len(score)
	scoreMap := make(map[int]int)
	for i := 0; i < n; i++ {
		scoreMap[score[i]] = i
	}
	ans := make([]string, n)
	sort.Slice(score, func(i, j int) bool { return score[i] > score[j] })
	for i, _ := range score {
		if i < 3 {

			ans[scoreMap[score[i]]] = desc[i]
		} else {
			ans[scoreMap[score[i]]] = strconv.Itoa(i + 1)
		}
	}
	return ans
}
