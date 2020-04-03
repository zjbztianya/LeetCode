package problem0017

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	digits string
	ans    []string
}{

	// 可以有多个 testcase
}

func Test_letterCombinations(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, letterCombinations(tc.digits), "输入:%v", tc)
	}
}

func Benchmark_letterCombinations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			letterCombinations(tc.digits)
		}
	}
}
