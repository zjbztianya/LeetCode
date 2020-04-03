package problem0143

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head *ListNode
	ans
}{

	// 可以有多个 testcase
}

func Test_reorderList(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, reorderList(tc.head), "输入:%v", tc)
	}
}

func Benchmark_reorderList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reorderList(tc.head)
		}
	}
}
