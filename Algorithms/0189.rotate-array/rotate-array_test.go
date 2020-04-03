package problem0189

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
 k int
	ans 
}{



	// 可以有多个 testcase
}

func Test_rotate(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, rotate(tc.nums, tc.k), "输入:%v", tc)
	}
}

func Benchmark_rotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			rotate(tc.nums, tc.k)
		}
	}
}
