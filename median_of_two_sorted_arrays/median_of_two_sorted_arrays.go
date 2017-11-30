package median_of_two_sorted_arrays

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func findKth(a []int, b []int, n int, m int, k int) int {
	if n > m {
		return findKth(b, a, m, n, k)
	}
	if n == 0 {
		return b[k-1]
	}
	if k == 1 {
		return min(a[0], b[0])
	}
	na := min(n, k/2)
	nb := k - na
	if a[na-1] == b[nb-1] {
		return a[na-1]
	}
	if a[na-1] < b[nb-1] {
		return findKth(a[na:], b, n-na, m, k-na)
	} else {
		return findKth(a, b[nb:], n, m-nb, k-nb)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	ans := 0.0
	ans += float64(findKth(nums1, nums2, n, m, (n+m)/2+1))
	if (n+m)&1 == 0 {
		ans += float64(findKth(nums1, nums2, n, m, (n+m)/2))
		ans /= 2
	}
	return ans
}