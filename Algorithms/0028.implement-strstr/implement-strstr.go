package problem0028

func getFail(s string) []int {
	f := make([]int, len(s))
	j := -1
	f[0] = -1
	for i := 1; i < len(s); i++ {
		for j != -1 && s[j+1] != s[i] {
			j = f[j]
		}
		if s[j+1] == s[i] {
			j++
		}
		f[i] = j
	}
	return f
}

func KMP(s, t string) int {
	if len(t) == 0 {
		return 0
	}

	f := getFail(t)
	j := -1
	for i := 0; i < len(s); i++ {
		for j != -1 && s[i] != t[j+1] {
			j = f[j]
		}
		if s[i] == t[j+1] {
			j++
		}
		if j+1 == len(t) {
			return i - len(t) + 1
		}
	}
	return -1
}

func strStr(haystack string, needle string) int {
	return KMP(haystack, needle)
}
