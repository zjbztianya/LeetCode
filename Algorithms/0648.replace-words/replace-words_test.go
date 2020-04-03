package problem0648

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	dict []string
 sentence string
	ans string
}{



	// 可以有多个 testcase
}

func Test_replaceWords(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, replaceWords(tc.dict, tc.sentence), "输入:%v", tc)
	}
}

func Benchmark_replaceWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			replaceWords(tc.dict, tc.sentence)
		}
	}
}
