package problem0215

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
 k int
	ans int
}{



	// 可以有多个 testcase
}

func Test_findKthLargest(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, findKthLargest(tc.nums, tc.k), "输入:%v", tc)
	}
}

func Benchmark_findKthLargest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findKthLargest(tc.nums, tc.k)
		}
	}
}
