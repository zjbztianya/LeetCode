package problem0207

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	numCourses int
 prerequisites [][]int
	ans bool
}{



	// 可以有多个 testcase
}

func Test_canFinish(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, canFinish(tc.numCourses, tc.prerequisites), "输入:%v", tc)
	}
}

func Benchmark_canFinish(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			canFinish(tc.numCourses, tc.prerequisites)
		}
	}
}
