package problem0088

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums := make([]int, n+m)
	for i, j, cnt := 0, 0, 0; i < m || j < n; cnt++ {
		if i < m && j < n {
			if nums1[i] < nums2[j] {
				nums[cnt] = nums1[i]
				i++
			} else {
				nums[cnt] = nums2[j]
				j++
			}
		} else if i < m {
			nums[cnt] = nums1[i]
			i++
		} else {
			nums[cnt] = nums2[j]
			j++
		}
	}
	copy(nums1, nums)
}
