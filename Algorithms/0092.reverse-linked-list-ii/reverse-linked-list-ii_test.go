package problem0092

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head *ListNode
	m    int
	n    int
	ans  *ListNode
}{

	// 可以有多个 testcase
}

func Test_reverseBetween(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, reverseBetween(tc.head, tc.m, tc.n), "输入:%v", tc)
	}
}

func Benchmark_reverseBetween(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reverseBetween(tc.head, tc.m, tc.n)
		}
	}
}
