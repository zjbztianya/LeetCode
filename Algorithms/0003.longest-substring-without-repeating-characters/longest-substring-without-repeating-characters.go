package problem0003

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	mp := make(map[uint8]int)
	ans := 0
	last_pos := -1
	for i := 0; i < len(s); i++ {
		if pos, ok := mp[s[i]]; ok && last_pos < pos {
			last_pos = pos
		}
		ans = max(ans, i-last_pos)
		mp[s[i]] = i
	}
	return ans
}