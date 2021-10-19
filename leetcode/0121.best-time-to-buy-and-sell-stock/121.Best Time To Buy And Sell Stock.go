package leetcode

/*会超出时间限制*/
func maxProfit(prices []int) int {
	maxprofit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if profit := prices[j] - prices[i]; profit > maxprofit {
				maxprofit = profit
			}
		}
	}
	return maxprofit
}

func maxProfit1(prices []int) int {
	inf := int(1e9)
	minprice := inf
	maxprofit := 0
	for _, v := range prices {
		if profit := v - minprice; profit > maxprofit {
			maxprofit = profit
		}
		if v < minprice {
			minprice = v
		}
	}
	return maxprofit
}
