package problem0027

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	val  int
	ans  int
}{

	// 可以有多个 testcase
}

func Test_removeElement(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, removeElement(tc.nums, tc.val), "输入:%v", tc)
	}
}

func Benchmark_removeElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			removeElement(tc.nums, tc.val)
		}
	}
}
