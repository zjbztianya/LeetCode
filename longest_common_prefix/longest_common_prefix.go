package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			if len(strs[j]) == i || strs[0][i] != strs[j][i] {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}
