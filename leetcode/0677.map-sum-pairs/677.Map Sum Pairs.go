package leetcode

import "strings"

type MapSum map[string]int

func Constructor() MapSum {
	return MapSum{}
}

func (m MapSum) Insert(key string, val int) {
	m[key] = val
}

func (m MapSum) Sum(prefix string) (sum int) {
	for s, v := range m {
		if strings.HasPrefix(s, prefix) {
			sum += v
		}
	}
	return
}

// type MapSum1 struct {
// 	m, pre map[string]int
// }

// func Constructor1() MapSum1 {
// 	return MapSum1{map[string]int{}, map[string]int{}}
// }

// func (m *MapSum1) Insert1(key string, val int) {

// }

// func (m *MapSum1) Sum1(prefix string) (sum int) {

// }

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
