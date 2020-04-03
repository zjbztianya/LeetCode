package problem0166

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	numerator   int
	denominator int
	ans         string
}{

	// 可以有多个 testcase
}

func Test_fractionToDecimal(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, fractionToDecimal(tc.numerator, tc.denominator), "输入:%v", tc)
	}
}

func Benchmark_fractionToDecimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			fractionToDecimal(tc.numerator, tc.denominator)
		}
	}
}
