package problem0088

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
	ans
}{

	// 可以有多个 testcase
}

func Test_merge(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, merge(tc.nums1, tc.m, tc.nums2, tc.n), "输入:%v", tc)
	}
}

func Benchmark_merge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			merge(tc.nums1, tc.m, tc.nums2, tc.n)
		}
	}
}
