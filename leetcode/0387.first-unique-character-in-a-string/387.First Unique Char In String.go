package leetcode

/*遍历两次数组 使用哈希表存储字符次数*/
func firstUniqChar(s string) int {
	charTimes := make(map[string]int)
	for _, v := range s {
		charTimes[string(v)]++
	}
	for k, v := range s {
		if charTimes[string(v)] <= 1 {
			return k
		}
	}
	return -1
}

type pair struct {
	ch  byte
	pos int
}

func firstUniqChar1(s string) int {
	m := len(s)
	pos := [25]int{}
	for i := range pos {
		pos[i] = m
	}
	q := []pair{}
	for i := range s {
		ch := s[i] - 'a'
		if pos[ch] == m {
			pos[ch] = i //将索引赋值
			q = append(q, pair{ch, i})
		} else {
			pos[ch] = m + 1
			/*队列先进先出*/
			for len(q) > 0 && pos[q[0].ch] == m+1 {
				q = q[1:]
			}
		}
	}
	if len(q) > 0 {
		return q[0].pos
	}
	return -1
}
