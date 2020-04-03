package problem0209

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s int
 nums []int
	ans int
}{



	// 可以有多个 testcase
}

func Test_minSubArrayLen(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, minSubArrayLen(tc.s, tc.nums), "输入:%v", tc)
	}
}

func Benchmark_minSubArrayLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minSubArrayLen(tc.s, tc.nums)
		}
	}
}
