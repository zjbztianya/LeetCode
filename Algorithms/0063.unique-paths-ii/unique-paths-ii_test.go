package problem0063

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	obstacleGrid [][]int
	ans          int
}{

	// 可以有多个 testcase
}

func Test_uniquePathsWithObstacles(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, uniquePathsWithObstacles(tc.obstacleGrid), "输入:%v", tc)
	}
}

func Benchmark_uniquePathsWithObstacles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			uniquePathsWithObstacles(tc.obstacleGrid)
		}
	}
}
