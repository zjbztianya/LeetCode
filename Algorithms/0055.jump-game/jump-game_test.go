package problem0055

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  bool
}{

	// 可以有多个 testcase
}

func Test_canJump(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, canJump(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_canJump(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			canJump(tc.nums)
		}
	}
}
