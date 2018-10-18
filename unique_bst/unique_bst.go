package unique_bst

func numTrees(n int) int {
	ans := 1
	for i := 1; i <= n; i++ {
		ans = ans * (2*n - i + 1) / i
	}
	return ans / (n + 1)
}
