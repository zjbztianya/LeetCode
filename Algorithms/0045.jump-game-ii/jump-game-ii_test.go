package problem0045

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  int
}{

	// 可以有多个 testcase
}

func Test_jump(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, jump(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_jump(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			jump(tc.nums)
		}
	}
}
