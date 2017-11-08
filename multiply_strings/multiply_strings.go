package multiply_strings

func multiply(num1 string, num2 string) string {
	n, m := len(num1), len(num2)
	ans := make([]int, n+m)
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			ans[n-i-1+m-j-1] += int(num1[i]-'0') * int(num2[j]-'0')
			ans[n-i-1+m-j] += ans[n-i-1+m-j-1] / 10
			ans[n-i-1+m-j-1] %= 10
		}
	}

	zlen := n + m
	for ; zlen > 1 && ans[zlen-1] == 0; zlen-- {
	}

	var ret []rune
	for i := zlen - 1; i >= 0; i-- {
		ret = append(ret, rune(ans[i]+'0'))
	}
	return string(ret)
}
