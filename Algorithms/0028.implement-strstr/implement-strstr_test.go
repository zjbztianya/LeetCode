package problem0028

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	haystack string
	needle   string
	ans      int
}{

	// 可以有多个 testcase
}

func Test_strStr(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, strStr(tc.haystack, tc.needle), "输入:%v", tc)
	}
}

func Benchmark_strStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			strStr(tc.haystack, tc.needle)
		}
	}
}
