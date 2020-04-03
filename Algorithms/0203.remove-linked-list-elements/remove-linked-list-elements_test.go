package problem0203

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head *ListNode
 val int
	ans *ListNode
}{



	// 可以有多个 testcase
}

func Test_removeElements(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, removeElements(tc.head, tc.val), "输入:%v", tc)
	}
}

func Benchmark_removeElements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			removeElements(tc.head, tc.val)
		}
	}
}
