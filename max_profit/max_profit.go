package max_profit

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	var ans int
	maxValue := prices[n-1]
	for i := n - 1; i >= 0; i-- {
		ans = max(ans, maxValue-prices[i])
		maxValue = max(maxValue, prices[i])
	}
	return ans
}
