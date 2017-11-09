package longest_substring_without_repeating_characters

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	mp := make(map[uint8]int)
	ans:=0
	last_pos := -1
	for i:=0; i < len(s); i++ {
		if _, ok:=mp[s[i]]; ok {
			ans = max(ans, i - last_pos)
			last_pos = i
		}else {
			mp[s[i]] = i
		}
	}
	ans = max(ans, len(s) - last_pos)
	return ans
}
