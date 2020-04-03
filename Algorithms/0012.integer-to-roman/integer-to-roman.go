package problem0012

func intToRoman(num int) string {
	var str string
	for num >= 1000 {
		str += "M"
		num -= 1000
	}
	if num/100 == 9 {
		str += "CM"
		num -= 900
	}
	if num >= 500 {
		str += "D"
		num -= 500
	}
	if num/100 == 4 {
		str += "CD"
		num -= 400
	}
	for num >= 100 {
		str += "C"
		num -= 100
	}
	if num/10 == 9 {
		str += "XC"
		num -= 90
	}
	if num >= 50 {
		str += "L"
		num -= 50
	}
	if num/10 == 4 {
		str += "XL"
		num -= 40
	}
	for num >= 10 {
		str += "X"
		num -= 10
	}
	if num == 9 {
		str += "IX"
		num = 0
	}
	if num >= 5 {
		str += "V"
		num -= 5
	}
	if num == 4 {
		str += "IV"
		num = 0
	}
	for num > 0 {
		str += "I"
		num--
	}
	return str
}
