package problem0166

import (
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	var b strings.Builder
	var sa, sb int
	if numerator < 0 {
		sa = 1
		numerator = -numerator
	}
	if denominator < 0 {
		sb = 1
		denominator = -denominator
	}
	sgn := sa ^ sb
	if sgn > 0 {
		b.WriteByte('-')
	}
	b.WriteString(strconv.Itoa(numerator / denominator))
	r := make(map[int]int)
	var a []byte
	writeFunc := func(pos int) {
		if pos == -1 {
			if len(a) > 0 {
				b.WriteByte('.')
			}
			b.Write(a)
			return
		}
		b.WriteByte('.')
		b.Write(a[:pos])
		b.WriteByte('(')
		b.Write(a[pos:])
		b.WriteByte(')')
	}
	for i := 1; ; i++ {
		m := numerator % denominator
		if r[m] > 0 {
			writeFunc(r[m] - 1)
			break
		} else {
			r[m] = i
			numerator = m * 10
			if m > 0 {
				a = append(a, byte(numerator/denominator+'0'))
			}
		}
		if m == 0 {
			writeFunc(-1)
			break
		}
	}
	return b.String()
}
