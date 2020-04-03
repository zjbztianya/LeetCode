package problem0013

func romanToInt(s string) int {
	n := len(s)
	var i, num int
	for i < n && s[i] == 'M' {
		num += 1000
		i++
	}
	if i+1 < n && s[i] == 'C' && s[i+1] == 'M' {
		num += 900
		i += 2
	}
	if i < n && s[i] == 'D' {
		num += 500
		i++
	}
	if i+1 < n && s[i] == 'C' && s[i+1] == 'D' {
		num += 400
		i += 2
	}
	for i < n && s[i] == 'C' {
		num += 100
		i++
	}
	if i+1 < n && s[i] == 'X' && s[i+1] == 'C' {
		num += 90
		i += 2
	}
	if i < n && s[i] == 'L' {
		num += 50
		i++
	}
	if i+1 < n && s[i] == 'X' && s[i+1] == 'L' {
		num += 40
		i += 2
	}
	for i < n && s[i] == 'X' {
		num += 10
		i++
	}
	if i+1 < n && s[i] == 'I' && s[i+1] == 'X' {
		num += 9
		i += 2
	}
	if i < n && s[i] == 'V' {
		num += 5
		i++
	}
	if i+1 < n && s[i] == 'I' && s[i+1] == 'V' {
		num += 4
		i += 2
	}
	for i < n && s[i] == 'I' {
		num++
		i++
	}
	return num
}
