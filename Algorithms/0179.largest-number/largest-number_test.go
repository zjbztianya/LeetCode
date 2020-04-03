package problem0179

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans string
}{



	// 可以有多个 testcase
}

func Test_largestNumber(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, largestNumber(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_largestNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			largestNumber(tc.nums)
		}
	}
}
