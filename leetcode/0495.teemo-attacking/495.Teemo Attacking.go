package leetcode

func findPoisonedDuration(timeSeries []int, duration int) int {
	res := duration
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i]-timeSeries[i-1] < duration {
			res += timeSeries[i] - timeSeries[i-1]
		} else {
			res += duration
		}
	}
	return res
}
