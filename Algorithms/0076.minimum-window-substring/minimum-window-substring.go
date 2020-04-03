package problem0076

import "math"

func minWindow(s string, t string) string {
	start, end, d, head := 0, 0, math.MaxInt32, -1
	count := len(t)
	mp := make([]int, 128)
	for i := 0; i < count; i++ {
		mp[t[i]]++
	}

	for end < len(s) {
		mp[s[end]]--
		if mp[s[end]] >= 0 {
			count--
		}
		end++
		for count == 0 {
			if end-start < d {
				d = end - start
				head = start
			}
			mp[s[start]]++
			if mp[s[start]] > 0 {
				count++
			}
			start++
		}
	}

	if head == -1 {
		return ""
	} else {
		return s[head : head+d]
	}
}
