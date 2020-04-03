package problem0004

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums1 []int
	nums2 []int
	ans   float64
}{

	// 可以有多个 testcase
}

func Test_findMedianSortedArrays(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, findMedianSortedArrays(tc.nums1, tc.nums2), "输入:%v", tc)
	}
}

func Benchmark_findMedianSortedArrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findMedianSortedArrays(tc.nums1, tc.nums2)
		}
	}
}
