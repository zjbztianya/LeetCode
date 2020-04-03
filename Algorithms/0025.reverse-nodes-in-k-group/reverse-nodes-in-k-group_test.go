package problem0025

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head *ListNode
	k    int
	ans  *ListNode
}{

	// 可以有多个 testcase
}

func Test_reverseKGroup(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, reverseKGroup(tc.head, tc.k), "输入:%v", tc)
	}
}

func Benchmark_reverseKGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reverseKGroup(tc.head, tc.k)
		}
	}
}
