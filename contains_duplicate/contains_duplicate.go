package contains_duplicate

func containsDuplicate(nums []int) bool {
	dup := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := dup[v]; ok {
			return true
		}
		dup[v] = struct{}{}
	}
	return false
}
