package set_mismatch

func findErrorNums(nums []int) []int {
	sum := 0
	var ans []int
	mp := make([]bool, len(nums))
	n := len(nums)
	var dupNum int
	for _, num := range nums {
		sum += num
		if !mp[num-1] {
			mp[num-1] = true
		} else {
			ans = append(ans, num)
			dupNum = num
		}
	}
	ans = append(ans, n*(n+1)/2-sum+dupNum)
	return ans
}
