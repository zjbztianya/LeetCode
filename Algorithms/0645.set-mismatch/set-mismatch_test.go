package problem0645

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans []int
}{



	// 可以有多个 testcase
}

func Test_findErrorNums(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, findErrorNums(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_findErrorNums(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findErrorNums(tc.nums)
		}
	}
}
