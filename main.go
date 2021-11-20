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
	integerReplacement(2020)
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
