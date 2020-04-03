package problem0024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head *ListNode
	ans  *ListNode
}{

	// 可以有多个 testcase
}

func Test_swapPairs(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, swapPairs(tc.head), "输入:%v", tc)
	}
}

func Benchmark_swapPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			swapPairs(tc.head)
		}
	}
}
