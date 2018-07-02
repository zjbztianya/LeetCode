package add_binary

func addBinary(a string, b string) string {
	n, m := len(a), len(b)
	if n > m {
		n, m = m, n
		a, b = b, a
	}
	ans := make([]byte, m+1)
	for i := m - 1; i >= 0; i-- {
		ans[i+1] += b[i] - '0'
		if i >= m-n {
			ans[i+1] += a[i+n-m] - '0'
		}
		ans[i] = ans[i+1] >> 1
		ans[i+1] = (ans[i+1] & 1) + '0'
	}
	ans[0] = ans[0] + '0'
	if ans[0] == '0' {
		return string(ans[1:])
	} else {
		return string(ans)
	}
}
