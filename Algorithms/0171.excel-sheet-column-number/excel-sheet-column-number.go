package problem0171
//171. Excel Sheet Column Number

func titleToNumber(s string) int {
	var res int
	for _, c := range s {
		res = res*26 + int(c-'A'+1)
	}
	return res
}

