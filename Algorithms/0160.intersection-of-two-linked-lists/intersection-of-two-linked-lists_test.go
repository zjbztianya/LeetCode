package problem0160

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	headA
	headB *ListNode
	ans   *ListNode
}{

	// 可以有多个 testcase
}

func Test_getIntersectionNode(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, getIntersectionNode(tc.headA, tc.headB), "输入:%v", tc)
	}
}

func Benchmark_getIntersectionNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			getIntersectionNode(tc.headA, tc.headB)
		}
	}
}
