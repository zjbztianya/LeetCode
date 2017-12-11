package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}
	var y int
	num := x
	for num > 0 {
		y = y*10 + num%10
		num /= 10
	}
	return x == y
}
