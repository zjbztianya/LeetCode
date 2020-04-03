package problem0023

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	lists []*ListNode
	ans   *ListNode
}{

	// 可以有多个 testcase
}

func Test_mergeKLists(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, mergeKLists(tc.lists), "输入:%v", tc)
	}
}

func Benchmark_mergeKLists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			mergeKLists(tc.lists)
		}
	}
}
