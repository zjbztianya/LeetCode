package palindromic_substrings

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Manacher(s string) int {
	ans := 0
	str := []byte("$#")
	for i := 0; i < len(s); i++ {
		str = append(str, s[i])
		str = append(str, '#')
	}

	mx, p := 0, 0
	n := len(str)
	r := make([]int, n)
	for i := range str {
		r[i] = 1
		if mx > i {
			r[i] = min(mx-i, r[2*p-i])
		}

		for i-r[i] >= 0 && i+r[i] < n && str[i-r[i]] == str[i+r[i]] {
			r[i]++
		}
		ans += r[i] >> 1

		if i+r[i] > mx {
			mx = i + r[i]
			p = i
		}
	}
	return ans
}
func countSubstrings(s string) int {
	return Manacher(s)
}
