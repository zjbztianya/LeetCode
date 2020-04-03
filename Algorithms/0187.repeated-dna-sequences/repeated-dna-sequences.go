package problem0187

func findRepeatedDnaSequences(s string) []string {
	couter := make(map[string]int)
	var ans []string
	n := len(s) - 9
	for i := 0; i < n; i++ {
		str := s[i : i+10]
		if cnt, ok := couter[str]; ok && cnt == 1 {
			ans = append(ans, str)
		}
		couter[str]++
	}
	return ans
}
