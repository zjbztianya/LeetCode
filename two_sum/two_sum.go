package two_sum

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))
	for i := range nums {
		if pos, ok := mp[target-nums[i]]; ok {
			return []int{pos, i}
		}
		mp[nums[i]] = i
	}
	return []int{}
}
