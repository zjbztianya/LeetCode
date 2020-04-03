package problem0647

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s string
	ans int
}{



	// 可以有多个 testcase
}

func Test_countSubstrings(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, countSubstrings(tc.s), "输入:%v", tc)
	}
}

func Benchmark_countSubstrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			countSubstrings(tc.s)
		}
	}
}
