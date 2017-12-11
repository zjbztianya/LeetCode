package letter_combinations

type Ans []string

func (ans *Ans) dfs(cur int, s []byte, digits string, kp []string) {
	if cur == len(digits) {
		*ans = append(*ans, string(s))
		return
	}
	str := kp[digits[cur]-'0']
	for i := 0; i < len(str); i++ {
		s = append(s, str[i])
		ans.dfs(cur+1, s, digits, kp)
		s = s[:cur]
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	kp := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var ans Ans
	ans.dfs(0, []byte{}, digits, kp)
	return ans
}
