package majority_element

//169. Majority Element

func majorityElement(nums []int) int {
	var c, m int
	for _, v := range nums {
		if c == 0 {
			m = v
		}
		if v == m {
			c++
		} else {
			c--
		}
	}
	return m
}
