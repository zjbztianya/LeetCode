package problem0018

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums   []int
	target int
	ans    [][]int
}{

	// 可以有多个 testcase
}

func Test_fourSum(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, fourSum(tc.nums, tc.target), "输入:%v", tc)
	}
}

func Benchmark_fourSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			fourSum(tc.nums, tc.target)
		}
	}
}
