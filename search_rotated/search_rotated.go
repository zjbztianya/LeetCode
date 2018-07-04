//package search_rotated
package main

import "fmt"

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	var mid int
	for l <= r {
		mid = (l + r) >> 1
		if nums[mid] == target {
			return true
		}
		if nums[mid] < nums[r] {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else if nums[mid] > nums[r] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			r--
		}
	}
	return false
}

func main() {
	var nums = []int{1, 1, 3, 1}
	fmt.Println(search(nums, 3))
}
