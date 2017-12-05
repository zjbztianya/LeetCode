package longest_palindromic_substring

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func Manacher(s string) (ra int, idx int) {
	var str []byte
	r := make([]int, len(s)*2+5)
	str = append(str, '$')
	str = append(str, '#')
	for i := range s {
		str = append(str, s[i])
		str = append(str, '#')
	}
	str = append(str, '@')
	ra, idx = 0, 0
	mx, p := 0, 0
	for i := 1; i < len(str)-1; i++ {
		if r[i] = 1; mx > i {
			r[i] = min(mx-i, r[2*p-i])
		}
		for {
			if str[i+r[i]] != str[i-r[i]] {
				break
			}
			r[i]++
		}
		if i+r[i] > mx {
			mx, p = i+r[i], i
		}
		if r[i] > ra {
			ra = r[i] - 1
			idx = i
		}
	}
	return
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	r, idx := Manacher(s)
	return s[(idx-r+1)/2-1 : (idx+r-1)/2]
}
