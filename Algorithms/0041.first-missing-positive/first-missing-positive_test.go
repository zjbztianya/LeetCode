package problem0041

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

func Test_firstMissingPositive(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, firstMissingPositive(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_firstMissingPositive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			firstMissingPositive(tc.nums)
		}
	}
}
