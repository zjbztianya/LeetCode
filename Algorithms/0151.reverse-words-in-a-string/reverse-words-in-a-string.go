package problem0151

import "bytes"

func reverse(r []byte) []byte {
	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-i-1] = r[n-i-1], r[i]
	}
	return r
}

func reverseWords(s string) string {
	r := bytes.TrimSpace([]byte(s))
	r = reverse(r)
	n, fieldStart := 0, 0
	for i := 0; i < len(r); i++ {
		if r[i] != ' ' {
			continue
		}
		n += copy(r[n:], reverse(r[fieldStart:i]))
		r[n] = ' '
		n++
		for i < len(r) && r[i] == ' ' {
			i++
		}
		fieldStart = i
	}
	if fieldStart < len(r) {
		n += copy(r[n:], reverse(r[fieldStart:]))
	}
	r = r[:n]
	return string(r)
}
