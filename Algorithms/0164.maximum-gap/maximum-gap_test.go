package problem0164

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

func Test_maximumGap(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, maximumGap(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_maximumGap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maximumGap(tc.nums)
		}
	}
}
