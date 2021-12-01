package leetcode

import "sort"

type fraction struct {
	x, y int
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	fractions := make([]*fraction, 0, n*(n-1)/2)
	for i, a := range arr {
		for _, b := range arr[i+1:] {
			fractions = append(fractions, &fraction{x: a, y: b})
		}
	}
	sort.Slice(fractions, func(i, j int) bool {
		a, b := fractions[i], fractions[j]
		return a.x*b.y < b.x*a.y
	})
	return []int{fractions[k-1].x, fractions[k-1].y}
}
