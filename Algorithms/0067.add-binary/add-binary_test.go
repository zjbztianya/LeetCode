package problem0067

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	a   string
	b   string
	ans string
}{

	// 可以有多个 testcase
}

func Test_addBinary(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, addBinary(tc.a, tc.b), "输入:%v", tc)
	}
}

func Benchmark_addBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			addBinary(tc.a, tc.b)
		}
	}
}
