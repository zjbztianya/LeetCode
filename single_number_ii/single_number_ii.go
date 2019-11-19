//package single_number_ii
package main

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

func main() {
	var a = []int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}
	singleNumber(a)
}
