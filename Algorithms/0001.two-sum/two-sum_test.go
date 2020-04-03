package problem0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums   []int
	target int
	ans    []int
}{

	// 可以有多个 testcase
}

func Test_twoSum(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, twoSum(tc.nums, tc.target), "输入:%v", tc)
	}
}

func Benchmark_twoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			twoSum(tc.nums, tc.target)
		}
	}
}
