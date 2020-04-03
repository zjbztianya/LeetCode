package problem0021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	l1  *ListNode
	l2  *ListNode
	ans *ListNode
}{

	// 可以有多个 testcase
}

func Test_mergeTwoLists(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, mergeTwoLists(tc.l1, tc.l2), "输入:%v", tc)
	}
}

func Benchmark_mergeTwoLists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			mergeTwoLists(tc.l1, tc.l2)
		}
	}
}
