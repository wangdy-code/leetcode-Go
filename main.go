package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"sync"
)

var count int
var wg sync.WaitGroup
var rw sync.RWMutex
var data int

type myValue struct {
	val int
	mu  sync.RWMutex
}

var memo map[int]int

func _integerReplacement(n int) (res int) {
	if n == 1 {
		return 0
	}
	if res, ok := memo[n]; ok {
		return res
	}
	defer func() { memo[n] = res }()
	if n%2 == 0 {
		return 1 + _integerReplacement(n/2)
	}
	return 2 + min(_integerReplacement(n/2), _integerReplacement(n/2+1))
}

func integerReplacement(n int) (res int) {
	memo = map[int]int{}
	return _integerReplacement(n)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	// cadence := sync.NewCond(&sync.Mutex{})
	// go func() {
	// 	for range time.Tick(1 * time.Millisecond) {
	// 		cadence.Broadcast()
	// 	}
	// }()
	// takeStep := func() {
	// 	cadence.L.Lock()
	// 	cadence.Wait()
	// 	cadence.L.Unlock()
	// }
	// tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
	// 	fmt.Fprintf(out, " %v", dirName)
	// 	atomic.AddInt32(dir, 1)
	// 	takeStep()
	// 	if atomic.LoadInt32(dir) == 1 {
	// 		fmt.Fprint(out, ". Success!")
	// 		return true
	// 	}
	// 	takeStep()
	// 	atomic.AddInt32(dir, -1)
	// 	return false
	// }
	// var left, right int32
	// tryLeft := func(out *bytes.Buffer) bool {
	// 	return tryDir("left", &left, out)
	// }
	// tryRight := func(out *bytes.Buffer) bool {
	// 	return tryDir("right", &right, out)
	// }
	// walk := func(walking *sync.WaitGroup, name string) {
	// 	var out bytes.Buffer
	// 	defer func() {
	// 		fmt.Println(out.String())
	// 	}()
	// 	defer walking.Done()
	// 	fmt.Fprintf(&out, "%v is trying to scoot:", name)
	// 	for i := 0; i < 5; i++ {
	// 		if tryLeft(&out) || tryRight(&out) {
	// 			return
	// 		}
	// 	}
	// 	fmt.Fprintf(&out, "\n&v tosses her hands up in exasperation!", name)
	// }
	// var peopleInHallway sync.WaitGroup
	// peopleInHallway.Add(2)
	// go walk(&peopleInHallway, "Alice")
	// go walk(&peopleInHallway, "Barbara")
	// peopleInHallway.Wait()
	println(1 % 3)
}

type Solution struct {
	m, n, total int
	mp          map[int]int
}

func Constructor(m int, n int) Solution {
	return Solution{m, n, m * n, map[int]int{}}
}

func (s *Solution) Flip() (ans []int) {
	x := rand.Intn(s.total)
	s.total--
	if y, ok := s.mp[x]; ok { // 查找位置 x 对应的映射
		ans = []int{y / s.n, y % s.n}
	} else {
		ans = []int{x / s.n, x % s.n}
	}
	if y, ok := s.mp[s.total]; ok { // 将位置 x 对应的映射设置为位置 total 对应的映射
		s.mp[x] = y
	} else {
		s.mp[x] = s.total
	}
	return
}

func (s *Solution) Reset() {
	s.total = s.m * s.n
	s.mp = map[int]int{}
}

func lengthOfLongestSubstring_(s string) int {
	charMap := make(map[byte]int)
	maxLen := 0
	n := len(s)
	start, end := 0, 0
	for start < n {
		if start != 0 {
			delete(charMap, s[start])
		}
		for end < len(s) && charMap[s[end]] == 0 {
			charMap[s[end]]++
			end++
		}
		maxLen = max(maxLen, end-start+1)
		start += end + 1

	}
	return maxLen
}
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	findGenerateParenthesis(n, n, "", &res)
	return res
}
func lengthOfLongestSubstring(s string) int {

	maxLen := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if findSubstring(s[i:j+1]) && maxLen < j+1-i {
				maxLen = j - i
			}
		}
	}
	return maxLen
}
func findSubstring(s string) bool {
	charMap := make(map[byte]int)
	for _, c := range s {
		if charMap[byte(c)] > 0 {
			return false
		}
		charMap[byte(c)]++
	}
	return true
}
func findGenerateParenthesis(lindex, rindex int, str string, res *[]string) {
	if lindex == 0 && rindex == 0 {
		*res = append(*res, str)
		return
	}
	if lindex > 0 {
		findGenerateParenthesis(lindex-1, rindex, str+"(", res)
	}
	if rindex > 0 && lindex < rindex {
		findGenerateParenthesis(lindex, rindex-1, str+")", res)
	}
}
func climbStairs(n int) int {
	sqrt5 := math.Sqrt(5)
	pow1 := math.Pow((1+sqrt5)/2, float64(n))
	pow2 := math.Pow((1-sqrt5)/2, float64(n))
	return int(math.Round((pow1 - pow2) / sqrt5))
}
func maxArea(height []int) int {
	ans, left := 0, 0
	right := len(height) - 1
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		ans = max(ans, area)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func read(n int) {
	rw.RLock()
	fmt.Printf("读goroutine %d 正在读取...\n", n)

	v := count

	fmt.Printf("读goroutine %d 读取结束，值为：%d\n", n, v)
	wg.Done()
	rw.RUnlock()
}

func write(n int) {
	rw.Lock()
	fmt.Printf("写goroutine %d 正在写入...\n", n)
	v := rand.Intn(1000)

	count = v

	fmt.Printf("写goroutine %d 写入结束，新值为：%d\n", n, v)
	wg.Done()
	rw.Unlock()
}
func reverse(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	if x < 0 {
		temp := abs(x)
		str := strconv.FormatInt(int64(temp), 10)
		var temp1 string
		for i := len(str) - 1; i >= 0; i-- {
			if str[i] != '0' {
				temp1 = temp1 + string(str[i])
			}
		}
		res, _ := strconv.Atoi(temp1)
		return 0 - res
	} else {
		temp := abs(x)
		str := strconv.FormatInt(int64(temp), 10)
		var temp1 string
		for i := len(str) - 1; i >= 0; i-- {
			if str[i] != '0' {
				temp1 = temp1 + string(str[i])
			}
		}
		res, _ := strconv.Atoi(temp1)
		return res
	}
}
func abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}
