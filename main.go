package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int)
	var ret bool
	for i, v := range nums {
		if j, ok := mp[v]; ok {
			if i-j < k {
				fmt.Println(j, i)
				ret = true
				break
			}
		}
		mp[v] = i
	}
	return ret
}

func main()  {
	a:=[]int{1,2,3,1,2,3}
	containsNearbyDuplicate(a, 2)
}