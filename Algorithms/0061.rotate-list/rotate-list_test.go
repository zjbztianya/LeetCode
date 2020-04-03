package problem0061

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

func Test_rotateRight(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, rotateRight(tc.head, tc.k), "输入:%v", tc)
	}
}

func Benchmark_rotateRight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			rotateRight(tc.head, tc.k)
		}
	}
}
