package problem0012

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num int
	ans string
}{

	// 可以有多个 testcase
}

func Test_intToRoman(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, intToRoman(tc.num), "输入:%v", tc)
	}
}

func Benchmark_intToRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			intToRoman(tc.num)
		}
	}
}
