package problem0058

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans int
}{

	// 可以有多个 testcase
}

func Test_lengthOfLastWord(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, lengthOfLastWord(tc.s), "输入:%v", tc)
	}
}

func Benchmark_lengthOfLastWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			lengthOfLastWord(tc.s)
		}
	}
}
