package problem0073

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	matrix [][]int
	ans
}{

	// 可以有多个 testcase
}

func Test_setZeroes(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, setZeroes(tc.matrix), "输入:%v", tc)
	}
}

func Benchmark_setZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			setZeroes(tc.matrix)
		}
	}
}
