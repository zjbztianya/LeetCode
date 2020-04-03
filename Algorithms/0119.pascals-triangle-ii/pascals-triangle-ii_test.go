package problem0119

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	rowIndex int
	ans      []int
}{

	// 可以有多个 testcase
}

func Test_getRow(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, getRow(tc.rowIndex), "输入:%v", tc)
	}
}

func Benchmark_getRow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			getRow(tc.rowIndex)
		}
	}
}
