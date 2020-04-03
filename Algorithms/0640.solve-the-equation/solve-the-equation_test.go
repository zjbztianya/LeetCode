package problem0640

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	equation string
	ans string
}{



	// 可以有多个 testcase
}

func Test_solveEquation(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, solveEquation(tc.equation), "输入:%v", tc)
	}
}

func Benchmark_solveEquation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			solveEquation(tc.equation)
		}
	}
}
