package problem0217

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans bool
}{



	// 可以有多个 testcase
}

func Test_containsDuplicate(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, containsDuplicate(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_containsDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			containsDuplicate(tc.nums)
		}
	}
}
