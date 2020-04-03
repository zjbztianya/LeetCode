package problem0125

import "strings"

func isLower(ch uint8) bool {
	if 'a' <= ch && ch <= 'z' {
		return true
	}
	return false
}

func isDigit(ch uint8) bool {
	if '0' <= ch && ch <= '9' {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		for l <= r && !isLower(s[l]) && !isDigit(s[l]) {
			l++
		}
		for l <= r && !isLower(s[r]) && !isDigit(s[r]) {
			r--
		}
		if l >= r {
			return true
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
