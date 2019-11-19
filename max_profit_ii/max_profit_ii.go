package max_profit_ii

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	var ans int
	var diff int
	for i := 1; i < n; i++ {
		diff = prices[i] - prices[i-1]
		if diff > 0 {
			ans += diff
		}
	}
	return ans
}
