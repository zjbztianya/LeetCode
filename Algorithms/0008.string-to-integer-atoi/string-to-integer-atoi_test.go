package problem0008

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	str string
	ans int
}{

	// 可以有多个 testcase
}

func Test_myAtoi(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, myAtoi(tc.str), "输入:%v", tc)
	}
}

func Benchmark_myAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			myAtoi(tc.str)
		}
	}
}
