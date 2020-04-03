package problem0165

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	version1 string
	version2 string
	ans      int
}{

	// 可以有多个 testcase
}

func Test_compareVersion(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, compareVersion(tc.version1, tc.version2), "输入:%v", tc)
	}
}

func Benchmark_compareVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			compareVersion(tc.version1, tc.version2)
		}
	}
}
