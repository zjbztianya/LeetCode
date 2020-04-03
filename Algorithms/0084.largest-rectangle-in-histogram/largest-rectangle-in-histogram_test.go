package problem0084

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	heights []int
	ans     int
}{

	// 可以有多个 testcase
}

func Test_largestRectangleArea(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, largestRectangleArea(tc.heights), "输入:%v", tc)
	}
}

func Benchmark_largestRectangleArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			largestRectangleArea(tc.heights)
		}
	}
}
