package problem0097

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s1  string
	s2  string
	s3  string
	ans bool
}{

	// 可以有多个 testcase
}

func Test_isInterleave(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, isInterleave(tc.s1, tc.s2, tc.s3), "输入:%v", tc)
	}
}

func Benchmark_isInterleave(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isInterleave(tc.s1, tc.s2, tc.s3)
		}
	}
}
