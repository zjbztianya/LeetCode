package problem0141

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head *ListNode
	ans  bool
}{

	// 可以有多个 testcase
}

func Test_hasCycle(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, hasCycle(tc.head), "输入:%v", tc)
	}
}

func Benchmark_hasCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			hasCycle(tc.head)
		}
	}
}
