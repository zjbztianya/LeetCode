package problem0137

func singleNumber(nums []int) int {
	var ans int32
	for i := 0; i < 32; i++ {
		cnt := 0
		for _, num := range nums {
			if (num & (1 << uint(i))) > 0 {
				cnt++
			}
		}
		if cnt%3 == 1 {
			ans |= 1 << uint(i)
		}
	}

	return int(ans)
}