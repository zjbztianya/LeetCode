package problem0167

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	numbers []int
	target  int
	ans     []int
}{

	// 可以有多个 testcase
}

func Test_twoSum(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, twoSum(tc.numbers, tc.target), "输入:%v", tc)
	}
}

func Benchmark_twoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			twoSum(tc.numbers, tc.target)
		}
	}
}
