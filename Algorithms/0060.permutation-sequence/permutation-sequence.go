package problem0060

var fac = []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

func getPermutation(n int, k int) string {
	var vis int
	k--
	ans := make([]byte, n)
	for i := 0; i < n; i++ {
		t := k / fac[n-i-1]
		var j uint
		for j = 0; j < uint(n); j++ {
			if (vis & (1 << j)) == 0 {
				if t == 0 {
					break
				}
				t--
			}
		}
		vis |= 1 << j
		ans[i] = byte(j + 1 + '0')
		k %= fac[n-i-1]
	}
	return string(ans)
}
