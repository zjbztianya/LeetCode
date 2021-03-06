package problem0146

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	capacity int
	ans      LRUCache
}{

	// 可以有多个 testcase
}

func Test_Constructor(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, Constructor(tc.capacity), "输入:%v", tc)
	}
}

func Benchmark_Constructor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			Constructor(tc.capacity)
		}
	}
}
