package problem0168

//168. Excel Sheet Column Title

func Reverse(r []byte) {
	n := len(r)
	for i := 0; i < len(r)/2; i++ {
		r[i], r[n-i-1] = r[n-i-1], r[i]
	}
}

func convertToTitle(n int) string {
	var res []byte
	for n > 0 {
		res = append(res, byte((n-1)%26+'A'))
		n = (n - 1) / 26
	}
	Reverse(res)
	return string(res)
}
