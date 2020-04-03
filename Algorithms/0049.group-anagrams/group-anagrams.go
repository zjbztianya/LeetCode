package problem0049

import "sort"

func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	mp := map[string]int{}
	for i := 0; i < len(strs); i++ {
		s := []byte(strs[i])
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		if _, ok := mp[string(s)]; !ok {
			ans = append(ans, []string{strs[i]})
			mp[string(s)] = len(ans) - 1
		} else {
			ans[mp[string(s)]] = append(ans[mp[string(s)]], strs[i])
		}
	}
	return ans
}