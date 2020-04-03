package problem0081

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums   []int
	target int
	ans    bool
}{

	// 可以有多个 testcase
}

func Test_search(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, search(tc.nums, tc.target), "输入:%v", tc)
	}
}

func Benchmark_search(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			search(tc.nums, tc.target)
		}
	}
}
