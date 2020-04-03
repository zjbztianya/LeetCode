package problem0019

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head *ListNode
	n    int
	ans  *ListNode
}{

	// 可以有多个 testcase
}

func Test_removeNthFromEnd(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, removeNthFromEnd(tc.head, tc.n), "输入:%v", tc)
	}
}

func Benchmark_removeNthFromEnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			removeNthFromEnd(tc.head, tc.n)
		}
	}
}
