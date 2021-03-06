package problem0082

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

func Test_deleteDuplicates(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, deleteDuplicates(tc.head), "输入:%v", tc)
	}
}

func Benchmark_deleteDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			deleteDuplicates(tc.head)
		}
	}
}
