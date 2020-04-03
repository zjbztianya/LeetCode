package problem0075

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans
}{

	// 可以有多个 testcase
}

func Test_sortColors(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, sortColors(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_sortColors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			sortColors(tc.nums)
		}
	}
}
