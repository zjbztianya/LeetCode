package problem0009

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	x   int
	ans bool
}{

	// 可以有多个 testcase
}

func Test_isPalindrome(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, isPalindrome(tc.x), "输入:%v", tc)
	}
}

func Benchmark_isPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isPalindrome(tc.x)
		}
	}
}
