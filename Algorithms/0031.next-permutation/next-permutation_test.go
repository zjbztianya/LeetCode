package problem0031

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans
}{

	// 可以有多个 testcase
}

func Test_nextPermutation(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, nextPermutation(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_nextPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			nextPermutation(tc.nums)
		}
	}
}
