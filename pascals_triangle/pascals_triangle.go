package pascals_triangle

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	nums := make([][]int, numRows)
	nums[0] = []int{1}
	for i := 1; i < numRows; i++ {
		nums[i] = make([]int, i+1)
		nums[i][0] = 1
		nums[i][i] = 1
		for j := 1; j < i; j++ {
			nums[i][j] = nums[i-1][j-1] + nums[i-1][j]
		}
	}
	return nums
}
