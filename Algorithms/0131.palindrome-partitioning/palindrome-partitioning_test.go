package problem0131

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans [][]string
}{

	// 可以有多个 testcase
}

func Test_partition(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, partition(tc.s), "输入:%v", tc)
	}
}

func Benchmark_partition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			partition(tc.s)
		}
	}
}
