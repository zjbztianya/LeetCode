package problem0643

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
 k int
	ans float64
}{



	// 可以有多个 testcase
}

func Test_findMaxAverage(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, findMaxAverage(tc.nums, tc.k), "输入:%v", tc)
	}
}

func Benchmark_findMaxAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findMaxAverage(tc.nums, tc.k)
		}
	}
}
