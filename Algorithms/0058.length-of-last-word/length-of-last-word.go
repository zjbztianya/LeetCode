package problem0058

func lengthOfLastWord(s string) int {
	var ans int
	pos := len(s) - 1
	for pos >= 0 && s[pos] == ' ' {
		pos--
	}
	for i := pos; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		ans++
	}
	return ans
}
