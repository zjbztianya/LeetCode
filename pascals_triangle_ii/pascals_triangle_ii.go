package pascals_triangle_ii

func getRow(rowIndex int) []int {
	nums := make([][]int, 2)
	nums[0] = make([]int, rowIndex+1)
	nums[1] = make([]int, rowIndex+1)
	nums[0][0] = 1
	nums[1][0] = 1
	for i := 1; i <= rowIndex; i++ {
		nums[i&1][i] = 1
		for j := 1; j < i; j++ {
			nums[i&1][j] = nums[(i+1)&1][j-1] + nums[(i+1)&1][j]
		}
	}
	return nums[rowIndex&1]
}
