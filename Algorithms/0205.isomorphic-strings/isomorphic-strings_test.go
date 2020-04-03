package problem0205

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s string
 t string
	ans bool
}{



	// 可以有多个 testcase
}

func Test_isIsomorphic(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, isIsomorphic(tc.s, tc.t), "输入:%v", tc)
	}
}

func Benchmark_isIsomorphic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isIsomorphic(tc.s, tc.t)
		}
	}
}
