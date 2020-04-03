package problem0210

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	numCourses int
 prerequisites [][]int
	ans []int
}{



	// 可以有多个 testcase
}

func Test_findOrder(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, findOrder(tc.numCourses, tc.prerequisites), "输入:%v", tc)
	}
}

func Benchmark_findOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findOrder(tc.numCourses, tc.prerequisites)
		}
	}
}
