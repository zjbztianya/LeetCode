package problem0015

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  [][]int
}{

	// 可以有多个 testcase
}

func Test_threeSum(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, threeSum(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_threeSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			threeSum(tc.nums)
		}
	}
}
