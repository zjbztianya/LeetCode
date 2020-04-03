package problem0030

func findSubstring(s string, words []string) []int {
	mp := make(map[string]int)
	ts := make(map[string]int)
	for i := 0; i < len(words); i++ {
		mp[words[i]]++
	}

	var ans []int
	n := len(words[0])
	for i := 0; i < n; i++ {
		ts = map[string]int{}
		cnt := 0
		for j := i; j < len(s)-n+1; j += n {
			curStr := s[j : j+n]
			if _, ok := mp[curStr]; ok {
				pos := j - cnt*n
				for _, ok := ts[curStr]; ok && ts[curStr] >= mp[curStr]; pos += n {
					ts[s[pos:pos+n]]--
					cnt--
				}
				ts[curStr]++
				cnt++
				if cnt == len(words) {
					ans = append(ans, j-(cnt-1)*n)
					ts[s[j-(cnt-1)*n:j-(cnt-2)*n]]--
					cnt--
				}
			} else {
				ts = map[string]int{}
				cnt = 0
			}
		}
	}
	return ans
}
