package solve_the_equation

import (
	"fmt"
	"strings"
)

func solve(expr string) (int, int) {
	expr += "$"
	flag := false
	x, sum, num, sgn := 0, 0, 0, 1
	for _, ch := range expr {
		if ch >= '0' && ch <= '9' {
			flag = true
			num = num*10 + int(ch-'0')
		} else if ch == 'x' {
			if num > 0 {
				x += num * sgn
			} else if !flag {
				x += sgn
			}
			num, sgn, flag = 0, 1, false
		} else {
			sum += num * sgn
			num, sgn, flag = 0, 1, false
			if ch == '-' {
				sgn = -1
			}
		}
	}
	return x, sum
}

func solveEquation(equation string) string {
	pos := strings.Index(equation, "=")
	lx, lsum := solve(equation[0:pos])
	rx, rsum := solve(equation[pos+1:])
	x := lx - rx
	sum := rsum - lsum
	if x != 0 {
		return fmt.Sprintf("x=%d", sum/x)
	} else if sum != 0 {
		return "No solution"
	} else {
		return "Infinite solutions"
	}
}
