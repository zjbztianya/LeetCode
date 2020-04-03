package problem0079

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	board [][]byte
	word  string
	ans   bool
}{

	// 可以有多个 testcase
}

func Test_exist(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, exist(tc.board, tc.word), "输入:%v", tc)
	}
}

func Benchmark_exist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			exist(tc.board, tc.word)
		}
	}
}
