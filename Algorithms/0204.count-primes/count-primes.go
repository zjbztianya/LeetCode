package problem0204

func countPrimes(n int) int {
	var prime []int
	vis := make([]bool, n)
	for i := 2; i < n; i++ {
		if !vis[i] {
			prime = append(prime, i)
		}
		for j := 0; j < len(prime) && i*prime[j] < n; j++ {
			vis[i*prime[j]] = true
			if i%prime[j] == 0 {
				break
			}
		}
	}
	return len(prime)
}

