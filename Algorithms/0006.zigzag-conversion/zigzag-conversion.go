package problem0006

func convert(s string, numRows int) string {
	row := make([][]byte, numRows)
	i, n := 0, len(s)
	for i < n {
		for j := 0; j < numRows && i < n; j++ {
			row[j] = append(row[j], s[i])
			i++
		}
		for j := numRows - 2; j > 0 && i < n; j-- {
			row[j] = append(row[j], s[i])
			i++
		}
	}
	var ans string
	for j := 0; j < numRows; j++ {
		ans += string(row[j])
	}
	return ans
}
