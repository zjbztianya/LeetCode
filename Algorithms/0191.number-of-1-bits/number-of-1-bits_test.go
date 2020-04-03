package problem0191

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num uint32
	ans int
}{



	// 可以有多个 testcase
}

func Test_hammingWeight(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, hammingWeight(tc.num), "输入:%v", tc)
	}
}

func Benchmark_hammingWeight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			hammingWeight(tc.num)
		}
	}
}
